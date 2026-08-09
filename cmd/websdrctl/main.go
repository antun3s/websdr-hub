// websdrctl: valida o catálogo, gera a API estática e roda o health check.
//
//	websdrctl validate -data data/stations
//	websdrctl build    -data data/stations -out dist/v1/stations.json
//	websdrctl check    -data data/stations -prev dist/v1/status.json -out dist/v1/status.json
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/websdrdir/websdr-directory/internal/catalog"
	"github.com/websdrdir/websdr-directory/internal/check"
	"github.com/websdrdir/websdr-directory/internal/status"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var err error
	switch os.Args[1] {
	case "validate":
		err = cmdValidate(os.Args[2:])
	case "build":
		err = cmdBuild(os.Args[2:])
	case "check":
		err = cmdCheck(ctx, os.Args[2:])
	case "-h", "--help", "help":
		usage()
		return
	default:
		fmt.Fprintf(os.Stderr, "subcomando desconhecido: %s\n\n", os.Args[1])
		usage()
		os.Exit(2)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "erro: %v\n", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprint(os.Stderr, `websdrctl <subcomando> [flags]

  validate   valida os YAML do catálogo (usado no CI de Pull Request)
  build      gera dist/v1/stations.json a partir do catálogo
  check      consulta as estações e gera dist/v1/status.json
`)
}

func cmdValidate(args []string) error {
	fs := flag.NewFlagSet("validate", flag.ExitOnError)
	dir := fs.String("data", "data/stations", "diretório do catálogo")
	if err := fs.Parse(args); err != nil {
		return err
	}

	stations, errs := catalog.Load(*dir)
	for _, e := range errs {
		fmt.Fprintf(os.Stderr, "  ✗ %v\n", e)
	}
	if len(errs) > 0 {
		return fmt.Errorf("%d problema(s) em %d estação(ões)", len(errs), len(stations))
	}
	fmt.Printf("✓ %d estações válidas\n", len(stations))
	return nil
}

func cmdBuild(args []string) error {
	fs := flag.NewFlagSet("build", flag.ExitOnError)
	dir := fs.String("data", "data/stations", "diretório do catálogo")
	out := fs.String("out", "dist/v1/stations.json", "arquivo de saída")
	if err := fs.Parse(args); err != nil {
		return err
	}

	stations, errs := catalog.Load(*dir)
	if len(errs) > 0 {
		return fmt.Errorf("catálogo inválido: rode `websdrctl validate` (%d problemas)", len(errs))
	}
	if err := os.MkdirAll(filepath.Dir(*out), 0o755); err != nil {
		return err
	}

	payload := struct {
		Version     int               `json:"version"`
		GeneratedAt string            `json:"generated_at"`
		Count       int               `json:"count"`
		Stations    []catalog.Station `json:"stations"`
	}{1, time.Now().UTC().Format(time.RFC3339), len(stations), stations}

	if err := status.WriteJSON(*out, payload); err != nil {
		return err
	}
	fmt.Printf("✓ %s (%d estações)\n", *out, len(stations))

	if err := copyFile("web/index.html", filepath.Join(filepath.Dir(filepath.Dir(*out)), "index.html")); err != nil {
		return fmt.Errorf("erro ao copiar index.html: %w", err)
	}
	fmt.Println("✓ dist/index.html")

	return nil
}

func cmdCheck(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("check", flag.ExitOnError)
	dir := fs.String("data", "data/stations", "diretório do catálogo")
	out := fs.String("out", "dist/v1/status.json", "arquivo de saída")
	prev := fs.String("prev", "", "status anterior: caminho local ou URL http(s)")
	conc := fs.Int("concurrency", 10, "consultas simultâneas")
	timeout := fs.Duration("timeout", 10*time.Second, "timeout por tentativa")
	retry := fs.Duration("retry", 5*time.Second, "espera entre a 1ª e a 2ª tentativa")
	vantage := fs.String("vantage", "local", "identificador do ponto de observação")
	if err := fs.Parse(args); err != nil {
		return err
	}

	stations, errs := catalog.Load(*dir)
	if len(errs) > 0 {
		return fmt.Errorf("catálogo inválido: rode `websdrctl validate` (%d problemas)", len(errs))
	}
	if err := os.MkdirAll(filepath.Dir(*out), 0o755); err != nil {
		return err
	}

	checker := check.New(*timeout, *retry)
	results := check.RunAll(ctx, checker, stations, *conc)
	merged := status.Merge(status.LoadPrevious(*prev), results, *vantage, time.Now())

	if err := status.WriteJSON(*out, merged); err != nil {
		return err
	}

	var online int
	for _, r := range results {
		if r.Online {
			online++
			continue
		}
		fmt.Printf("  offline  %-28s %s\n", r.ID, r.Error)
	}
	fmt.Printf("✓ %s — %d/%d online\n", *out, online, len(results))

	// Exit 0 mesmo com estações offline: estação fora do ar é o dado que o
	// projeto existe para coletar, não uma falha do pipeline.
	return nil
}

func copyFile(src, dst string) error {
	s, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("abrindo origem: %w", err)
	}
	defer s.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	d, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("criando destino: %w", err)
	}
	defer d.Close()

	if _, err := io.Copy(d, s); err != nil {
		return fmt.Errorf("copiando: %w", err)
	}
	return d.Close()
}
