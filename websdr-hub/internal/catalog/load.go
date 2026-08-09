package catalog

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/goccy/go-yaml"
)

// Load lê todos os *.yaml do diretório, valida cada estação e depois
// aplica as regras que só existem no conjunto (id/URL duplicados).
// Devolve o catálogo ordenado por id e a lista completa de problemas.
func Load(dir string) ([]Station, []error) {
	var (
		stations []Station
		errs     []error
	)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, []error{fmt.Errorf("lendo %s: %w", dir, err)}
	}

	byID := map[string]string{}   // id -> arquivo
	byHost := map[string]string{} // host normalizado -> arquivo

	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || (!strings.HasSuffix(name, ".yaml") && !strings.HasSuffix(name, ".yml")) {
			continue
		}
		path := filepath.Join(dir, name)

		raw, err := os.ReadFile(path)
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", name, err))
			continue
		}

		var st Station
		if err := yaml.Unmarshal(raw, &st); err != nil {
			errs = append(errs, fmt.Errorf("%s: YAML inválido: %w", name, err))
			continue
		}

		for _, ve := range st.Validate() {
			errs = append(errs, fmt.Errorf("%s: %w", name, ve))
		}

		// Segundo parse, estrito: pega typo em chave (ex: "cordinates") que
		// de outro modo viraria silenciosamente zero value. Fica depois da
		// validação semântica de propósito — o contribuidor precisa ver
		// todos os problemas do arquivo numa única rodada de CI, não um por vez.
		if err := yaml.UnmarshalWithOptions(raw, &Station{}, yaml.DisallowUnknownField()); err != nil {
			errs = append(errs, fmt.Errorf("%s: campo não reconhecido: %w", name, err))
		}

		// O nome do arquivo espelha o id: torna a PR legível no diff e
		// impede dois contribuidores de criarem o mesmo receptor em arquivos
		// com nomes diferentes.
		if want := st.ID + ".yaml"; st.ID != "" && name != want {
			errs = append(errs, fmt.Errorf("%s: o arquivo deveria se chamar %s (igual ao id)", name, want))
		}

		if prev, dup := byID[st.ID]; dup && st.ID != "" {
			errs = append(errs, fmt.Errorf("%s: id %q já usado em %s", name, st.ID, prev))
		} else if st.ID != "" {
			byID[st.ID] = name
		}

		if key := normalizeURL(st.URL); key != "" {
			if prev, dup := byHost[key]; dup {
				errs = append(errs, fmt.Errorf("%s: URL %s duplica a de %s", name, st.URL, prev))
			} else {
				byHost[key] = name
			}
		}

		stations = append(stations, st)
	}

	sort.Slice(stations, func(i, j int) bool { return stations[i].ID < stations[j].ID })
	return stations, errs
}

// normalizeURL reduz a URL a host:porta em minúsculas, sem barra final.
// É o suficiente para detectar o caso real de duplicata: mesma estação
// cadastrada como http://x:8073 e http://x:8073/.
func normalizeURL(raw string) string {
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		return ""
	}
	return strings.ToLower(u.Host) + strings.TrimSuffix(u.Path, "/")
}
