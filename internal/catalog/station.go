// Package catalog carrega e valida o catálogo curado de estações.
package catalog

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// Software conhecido. Determina o probe usado no health check.
var knownSoftware = map[string]bool{
	"websdr":     true,
	"kiwisdr":    true,
	"openwebrx":  true,
	"phantomsdr": true,
	"other":      true,
}

var (
	reSlug    = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)
	reCountry = regexp.MustCompile(`^[A-Z]{2}$`)
)

// Station é a unidade curada do catálogo: um arquivo YAML por estação.
type Station struct {
	ID        string   `yaml:"id" json:"id"`
	Name      string   `yaml:"name" json:"name"`
	URL       string   `yaml:"url" json:"url"`
	Software  string   `yaml:"software" json:"software"`
	Location  Location `yaml:"location" json:"location"`
	Languages []string `yaml:"languages,omitempty" json:"languages,omitempty"`
	Coverage  []Band   `yaml:"coverage" json:"coverage"`
	MaxUsers  int      `yaml:"max_users,omitempty" json:"max_users,omitempty"`
	Operator  string   `yaml:"operator,omitempty" json:"operator,omitempty"` // apenas indicativo (callsign)
	AddedAt   string   `yaml:"added_at" json:"added_at"`
}

type Location struct {
	Country     string     `yaml:"country" json:"country"`         // ISO 3166-1 alpha-2
	City        string     `yaml:"city" json:"city"`               //
	Coordinates [2]float64 `yaml:"coordinates" json:"coordinates"` // [lat, lon]
}

// Band usa inteiros em Hz de propósito: float em MHz perde precisão no
// round-trip YAML->JSON e inviabiliza range query exata.
type Band struct {
	Name    string   `yaml:"name" json:"name"`
	StartHz int64    `yaml:"start_hz" json:"start_hz"`
	StopHz  int64    `yaml:"stop_hz" json:"stop_hz"`
	Modes   []string `yaml:"modes,omitempty" json:"modes,omitempty"`
}

// Validate devolve todos os problemas encontrados, não apenas o primeiro:
// o contribuidor corrige tudo numa única iteração da PR.
func (s *Station) Validate() []error {
	var errs []error
	bad := func(format string, a ...any) { errs = append(errs, fmt.Errorf(format, a...)) }

	if !reSlug.MatchString(s.ID) {
		bad("id %q: precisa ser um slug kebab-case ([a-z0-9-])", s.ID)
	}
	if strings.TrimSpace(s.Name) == "" {
		bad("name: obrigatório")
	}

	u, err := url.Parse(s.URL)
	switch {
	case err != nil:
		bad("url %q: inválida: %v", s.URL, err)
	case u.Scheme != "http" && u.Scheme != "https":
		bad("url %q: esquema precisa ser http ou https", s.URL)
	case u.Host == "":
		bad("url %q: host ausente", s.URL)
	}

	if !knownSoftware[s.Software] {
		bad("software %q: valores aceitos: websdr, kiwisdr, openwebrx, phantomsdr, other", s.Software)
	}

	if !reCountry.MatchString(s.Location.Country) {
		bad("location.country %q: use ISO 3166-1 alpha-2 maiúsculo (ex: BR, NL)", s.Location.Country)
	}
	if strings.TrimSpace(s.Location.City) == "" {
		bad("location.city: obrigatório")
	}

	lat, lon := s.Location.Coordinates[0], s.Location.Coordinates[1]
	switch {
	case lat < -90 || lat > 90:
		bad("location.coordinates: latitude %v fora de [-90, 90]", lat)
	case lon < -180 || lon > 180:
		bad("location.coordinates: longitude %v fora de [-180, 180]", lon)
	case lat == 0 && lon == 0:
		bad("location.coordinates: [0, 0] normalmente indica coordenada esquecida")
	}

	if len(s.Coverage) == 0 {
		bad("coverage: informe ao menos uma faixa")
	}
	for i, b := range s.Coverage {
		if strings.TrimSpace(b.Name) == "" {
			bad("coverage[%d].name: obrigatório", i)
		}
		if b.StartHz < 0 {
			bad("coverage[%d]: start_hz negativo", i)
		}
		if b.StopHz <= b.StartHz {
			bad("coverage[%d]: stop_hz (%d) precisa ser maior que start_hz (%d)", i, b.StopHz, b.StartHz)
		}
	}

	if _, err := time.Parse("2006-01-02", s.AddedAt); err != nil {
		bad("added_at %q: use o formato YYYY-MM-DD", s.AddedAt)
	}

	return errs
}
