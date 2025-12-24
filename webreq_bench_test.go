package webreq_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/tonnytg/webreq"
)

// BenchmarkExecute_Sequential measures sequential request performance
func BenchmarkExecute_Sequential(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("benchmark-response"))
	}))
	defer ts.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := webreq.NewRequest("GET")
		req.SetURL(ts.URL)
		req.SetTimeout(5)
		_, err := req.Execute()
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

// BenchmarkExecute_Parallel measures concurrent request performance
func BenchmarkExecute_Parallel(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("benchmark-response"))
	}))
	defer ts.Close()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := webreq.NewRequest("GET")
			req.SetURL(ts.URL)
			req.SetTimeout(5)
			_, err := req.Execute()
			if err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
		}
	})
}

// BenchmarkExecute_WithHeaders measures performance with headers
func BenchmarkExecute_WithHeaders(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("benchmark-response"))
	}))
	defer ts.Close()

	headers := webreq.HeadersMap{
		"Content-Type":  "application/json",
		"Authorization": "Bearer token123",
		"User-Agent":    "WebReq/1.0",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := webreq.NewRequest("GET")
		req.SetURL(ts.URL)
		req.SetHeaders(headers)
		req.SetTimeout(5)
		_, err := req.Execute()
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

// BenchmarkExecute_POST_WithBody measures POST performance with body
func BenchmarkExecute_POST_WithBody(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("benchmark-response"))
	}))
	defer ts.Close()

	data := []byte(`{"key":"value","name":"test","count":42}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := webreq.NewRequest("POST")
		req.SetURL(ts.URL)
		req.SetData(data)
		req.SetTimeout(5)
		_, err := req.Execute()
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

// BenchmarkExecute_HighConcurrency tests performance under high concurrent load
func BenchmarkExecute_HighConcurrency(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("benchmark-response"))
	}))
	defer ts.Close()

	const concurrency = 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(concurrency)
		for j := 0; j < concurrency; j++ {
			go func() {
				defer wg.Done()
				req := webreq.NewRequest("GET")
				req.SetURL(ts.URL)
				req.SetTimeout(5)
				_, _ = req.Execute()
			}()
		}
		wg.Wait()
	}
}

// BenchmarkExecuteWithContext measures performance with custom context
func BenchmarkExecuteWithContext(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("benchmark-response"))
	}))
	defer ts.Close()

	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := webreq.NewRequest("GET")
		req.SetURL(ts.URL)
		_, err := req.ExecuteWithContext(ctx)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
