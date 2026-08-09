// Package check faz o probe de disponibilidade das estações.
package check

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/websdrdir/websdr-directory/internal/catalog"
)

const (
	// Corpo lido e descartado: confirma que o servidor realmente entrega
	// dados, não apenas completa o handshake e trava.
	bodySniff = 4 << 10

	// Identificação honesta. Muito WebSDR roda em link residencial e o
	// operador precisa saber quem está batendo na porta dele.
	DefaultUserAgent = "websdr-directory/0.1 (+https://github.com/websdrdir/websdr-directory)"
)

// Result é o resultado consolidado das tentativas de uma estação.
type Result struct {
	ID        string
	Online    bool
	HTTPCode  int
	LatencyMS int64
	Attempts  int
	CheckedAt time.Time
	Error     string
}

type attempt struct {
	ok        bool
	httpCode  int
	latencyMS int64
	err       string
}

// Checker é reutilizável entre estações; o http.Client compartilha o pool
// de conexões e é seguro para uso concorrente.
type Checker struct {
	Client    *http.Client
	UserAgent string
	Timeout   time.Duration // por tentativa
	Retry     time.Duration // espera entre a 1ª e a 2ª tentativa
	Attempts  int           // MVP: 2
	Now       func() time.Time
}

func New(timeout, retry time.Duration) *Checker {
	return &Checker{
		Client: &http.Client{
			// Sem seguir mais de 5 redirects: alguns receptores mortos
			// caem em loop de portal cativo do provedor.
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				if len(via) >= 5 {
					return http.ErrUseLastResponse
				}
				return nil
			},
		},
		UserAgent: DefaultUserAgent,
		Timeout:   timeout,
		Retry:     retry,
		Attempts:  2,
		Now:       time.Now,
	}
}

func (c *Checker) probe(ctx context.Context, rawURL string) attempt {
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return attempt{err: err.Error()}
	}
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", "*/*")

	start := c.Now()
	resp, err := c.Client.Do(req)
	if err != nil {
		return attempt{latencyMS: c.Now().Sub(start).Milliseconds(), err: err.Error()}
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, bodySniff))
	latency := c.Now().Sub(start).Milliseconds()

	a := attempt{httpCode: resp.StatusCode, latencyMS: latency}
	if resp.StatusCode >= 400 {
		a.err = fmt.Sprintf("HTTP %d", resp.StatusCode)
		return a
	}
	a.ok = true
	return a
}

// Check executa até duas consultas. A segunda existe para absorver o falso
// negativo isolado (timeout de link residencial, rate limit momentâneo);
// só marca offline quando as duas falham.
func (c *Checker) Check(ctx context.Context, st catalog.Station) Result {
	res := Result{ID: st.ID, CheckedAt: c.Now().UTC()}

	for i := 0; i < c.Attempts; i++ {
		if i > 0 {
			select {
			case <-ctx.Done():
				res.Error = ctx.Err().Error()
				return res
			case <-time.After(c.Retry):
			}
		}

		a := c.probe(ctx, st.URL)
		res.Attempts = i + 1
		res.HTTPCode, res.LatencyMS, res.Error = a.httpCode, a.latencyMS, a.err

		if a.ok {
			res.Online = true
			res.Error = ""
			return res
		}
	}
	return res
}

// RunAll respeita um teto de concorrência: o objetivo não é terminar rápido,
// é não parecer um scanner para centenas de Raspberry Pis simultaneamente.
func RunAll(ctx context.Context, c *Checker, stations []catalog.Station, concurrency int) []Result {
	if concurrency < 1 {
		concurrency = 1
	}
	results := make([]Result, len(stations))
	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for i, st := range stations {
		wg.Add(1)
		go func(i int, st catalog.Station) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			results[i] = c.Check(ctx, st)
		}(i, st)
	}
	wg.Wait()
	return results
}
