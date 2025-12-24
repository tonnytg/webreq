package webreq_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/tonnytg/webreq"
)

func TestExecuteWithContext_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("context-success"))
	}))
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := webreq.NewRequest("GET")
	req.SetURL(ts.URL)

	body, err := req.ExecuteWithContext(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != "context-success" {
		t.Fatalf("unexpected body: %q", string(body))
	}
}

func TestExecuteWithContext_Cancellation(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("too-late"))
	}))
	defer ts.Close()

	ctx, cancel := context.WithCancel(context.Background())
	
	req := webreq.NewRequest("GET")
	req.SetURL(ts.URL)

	// Cancel immediately
	cancel()

	_, err := req.ExecuteWithContext(ctx)
	if err == nil {
		t.Fatal("expected error due to cancelled context, got nil")
	}
}

func TestExecuteWithContext_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("too-late"))
	}))
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	req := webreq.NewRequest("GET")
	req.SetURL(ts.URL)

	_, err := req.ExecuteWithContext(ctx)
	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}
}

func TestExecuteWithContext_WithHeaders(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Custom") != "test-value" {
			t.Errorf("expected X-Custom header, got %q", r.Header.Get("X-Custom"))
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}))
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := webreq.NewRequest("GET")
	req.SetURL(ts.URL)
	req.SetHeaders(webreq.HeadersMap{"X-Custom": "test-value"})

	body, err := req.ExecuteWithContext(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != "ok" {
		t.Fatalf("unexpected body: %q", string(body))
	}
}

func TestExecuteWithContext_POST(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("created"))
	}))
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := webreq.NewRequest("POST")
	req.SetURL(ts.URL)
	req.SetData([]byte(`{"test":"data"}`))

	body, err := req.ExecuteWithContext(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != "created" {
		t.Fatalf("unexpected body: %q", string(body))
	}
}
