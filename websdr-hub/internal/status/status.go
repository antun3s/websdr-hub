// Package status monta o status.json público a partir dos resultados da
// rodada atual e do estado da rodada anterior.
package status

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/websdrdir/websdr-directory/internal/check"
)

type Entry struct {
	State               string `json:"state"` // online | offline
	HTTPCode            int    `json:"http_code,omitempty"`
	LatencyMS           int64  `json:"latency_ms,omitempty"`
	CheckedAt           string `json:"checked_at"`
	ConsecutiveFailures int    `json:"consecutive_failures"`
	LastOnline          string `json:"last_online,omitempty"`
	Error               string `json:"error,omitempty"`
}

type File struct {
	Version     int              `json:"version"`
	GeneratedAt string           `json:"generated_at"`
	Vantage     string           `json:"vantage"`
	Stations    map[string]Entry `json:"stations"`
}

// LoadPrevious aceita caminho local ou URL http(s). O estado anterior mora
// no próprio site publicado, então não há branch nem commit de estado —
// se a leitura falhar (primeira execução, Pages fora do ar), começa vazio.
func LoadPrevious(src string) File {
	empty := File{Version: 1, Stations: map[string]Entry{}}
	if src == "" {
		return empty
	}

	var raw []byte
	var err error

	if strings.HasPrefix(src, "http://") || strings.HasPrefix(src, "https://") {
		client := &http.Client{Timeout: 20 * time.Second}
		resp, rerr := client.Get(src)
		if rerr != nil {
			return empty
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return empty
		}
		raw, err = io.ReadAll(io.LimitReader(resp.Body, 16<<20))
	} else {
		raw, err = os.ReadFile(src)
	}
	if err != nil {
		return empty
	}

	var f File
	if json.Unmarshal(raw, &f) != nil || f.Stations == nil {
		return empty
	}
	return f
}

// Merge produz o novo status. consecutive_failures e last_online são a única
// memória que o MVP carrega entre execuções: dão o sinal para podar estações
// mortas sem precisar de série temporal ainda.
func Merge(prev File, results []check.Result, vantage string, now time.Time) File {
	out := File{
		Version:     1,
		GeneratedAt: now.UTC().Format(time.RFC3339),
		Vantage:     vantage,
		Stations:    make(map[string]Entry, len(results)),
	}

	for _, r := range results {
		old := prev.Stations[r.ID]
		e := Entry{
			HTTPCode:  r.HTTPCode,
			LatencyMS: r.LatencyMS,
			CheckedAt: r.CheckedAt.Format(time.RFC3339),
			Error:     r.Error,
		}
		if r.Online {
			e.State = "online"
			e.ConsecutiveFailures = 0
			e.LastOnline = e.CheckedAt
		} else {
			e.State = "offline"
			e.ConsecutiveFailures = old.ConsecutiveFailures + 1
			e.LastOnline = old.LastOnline
		}
		out.Stations[r.ID] = e
	}
	return out
}

// WriteJSON grava de forma atômica: o deploy nunca publica um arquivo
// truncado se o processo morrer no meio da escrita.
func WriteJSON(path string, v any) error {
	raw, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	raw = append(raw, '\n')

	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, raw, 0o644); err != nil {
		return fmt.Errorf("escrevendo %s: %w", tmp, err)
	}
	return os.Rename(tmp, path)
}
