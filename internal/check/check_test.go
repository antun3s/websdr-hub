package check

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/websdrdir/websdr-directory/internal/catalog"
)

func newChecker() *Checker {
	c := New(2*time.Second, time.Millisecond) // retry curto no teste
	return c
}

func TestOnlineNaPrimeiraTentativa(t *testing.T) {
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
		if got := r.Header.Get("User-Agent"); got != DefaultUserAgent {
			t.Errorf("User-Agent = %q, esperado %q", got, DefaultUserAgent)
		}
		w.Write([]byte("websdr"))
	}))
	defer srv.Close()

	res := newChecker().Check(context.Background(), catalog.Station{ID: "x", URL: srv.URL})

	if !res.Online || res.HTTPCode != 200 {
		t.Fatalf("esperado online/200, obtido %+v", res)
	}
	if res.Attempts != 1 {
		t.Errorf("Attempts = %d, esperado 1 (não deve repetir se a 1ª deu certo)", res.Attempts)
	}
	if n := atomic.LoadInt32(&hits); n != 1 {
		t.Errorf("servidor recebeu %d requisições, esperado 1", n)
	}
}

// O caso que justifica as duas consultas: falha isolada seguida de sucesso
// não pode marcar a estação como offline.
func TestSegundaTentativaRecupera(t *testing.T) {
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&hits, 1) == 1 {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.Write([]byte("ok"))
	}))
	defer srv.Close()

	res := newChecker().Check(context.Background(), catalog.Station{ID: "x", URL: srv.URL})

	if !res.Online {
		t.Fatalf("esperado online após retry, obtido %+v", res)
	}
	if res.Attempts != 2 {
		t.Errorf("Attempts = %d, esperado 2", res.Attempts)
	}
	if res.Error != "" {
		t.Errorf("Error = %q, deveria ser limpo após sucesso", res.Error)
	}
}

func TestOfflineAposDuasFalhas(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	res := newChecker().Check(context.Background(), catalog.Station{ID: "x", URL: srv.URL})

	if res.Online {
		t.Fatal("esperado offline")
	}
	if res.Attempts != 2 || res.HTTPCode != 503 || res.Error != "HTTP 503" {
		t.Errorf("resultado inesperado: %+v", res)
	}
}

func TestHostInexistente(t *testing.T) {
	res := newChecker().Check(context.Background(),
		catalog.Station{ID: "x", URL: "http://nao-existe.invalid:8073/"})

	if res.Online || res.Error == "" {
		t.Errorf("esperado offline com erro de rede, obtido %+v", res)
	}
}

func TestRunAllRespeitaConcorrencia(t *testing.T) {
	var inFlight, maxSeen int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cur := atomic.AddInt32(&inFlight, 1)
		for {
			old := atomic.LoadInt32(&maxSeen)
			if cur <= old || atomic.CompareAndSwapInt32(&maxSeen, old, cur) {
				break
			}
		}
		time.Sleep(20 * time.Millisecond)
		atomic.AddInt32(&inFlight, -1)
		w.Write([]byte("ok"))
	}))
	defer srv.Close()

	stations := make([]catalog.Station, 12)
	for i := range stations {
		stations[i] = catalog.Station{ID: string(rune('a' + i)), URL: srv.URL}
	}

	results := RunAll(context.Background(), newChecker(), stations, 3)

	if len(results) != 12 {
		t.Fatalf("len(results) = %d, esperado 12", len(results))
	}
	for i, r := range results {
		if !r.Online {
			t.Errorf("results[%d] offline: %+v", i, r)
		}
	}
	if got := atomic.LoadInt32(&maxSeen); got > 3 {
		t.Errorf("concorrência máxima observada = %d, teto era 3", got)
	}
}
