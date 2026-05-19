package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakeRequest_200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	result := make(chan int, 1)
	makeRequest(srv.URL, result)

	if got := <-result; got != http.StatusOK {
		t.Errorf("got %d, want %d", got, http.StatusOK)
	}
}

func TestMakeRequest_503(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	result := make(chan int, 1)
	makeRequest(srv.URL, result)

	if got := <-result; got != http.StatusServiceUnavailable {
		t.Errorf("got %d, want %d", got, http.StatusServiceUnavailable)
	}
}

func TestMakeRequest_ConnectionRefused(t *testing.T) {
	result := make(chan int, 1)
	makeRequest("http://localhost:1", result)

	if got := <-result; got != 500 {
		t.Errorf("got %d, want 500", got)
	}
}

func BenchmarkMakeRequest(b *testing.B) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	b.ResetTimer()
	for range b.N {
		result := make(chan int, 1)
		makeRequest(srv.URL, result)
		<-result
	}
}
