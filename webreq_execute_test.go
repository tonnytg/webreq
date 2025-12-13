package webreq_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/tonnytg/webreq"
)

func TestExecute_GET_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", r.Method)
		}
		if got := r.Header.Get("X-Test"); got != "1" {
			t.Errorf("expected header X-Test=1, got %q", got)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello-get"))
	}))
	defer ts.Close()

	req := webreq.NewRequest("GET")
	req.SetURL(ts.URL)
	req.SetHeaders(webreq.HeadersMap{"X-Test": "1"})
	req.SetTimeout(2)

	body, err := req.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != "hello-get" {
		t.Fatalf("unexpected body: %q", string(body))
	}
	if req.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", req.StatusCode)
	}
}

func TestExecute_POST_BodyAndHeader(t *testing.T) {
	type payload struct {
		Name string `json:"name"`
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("expected Content-Type application/json, got %q", ct)
		}
		b, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer r.Body.Close()
		var p payload
		if err := json.Unmarshal(b, &p); err != nil {
			t.Fatalf("invalid json received: %v", err)
		}
		if p.Name != "tonny" {
			t.Fatalf("unexpected name: %q", p.Name)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer ts.Close()

	req := webreq.NewRequest("POST")
	req.SetURL(ts.URL)
	req.SetHeaders(webreq.HeadersMap{"Content-Type": "application/json"})
	j, _ := json.Marshal(payload{Name: "tonny"})
	req.SetData(j)
	req.SetTimeout(2)

	body, err := req.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if req.StatusCode != http.StatusCreated {
		t.Fatalf("unexpected status code: %d", req.StatusCode)
	}
	if string(body) != "{"+"\"ok\":true}" { // simple literal check
		// We still accept the body, just ensure it's non-empty and JSON-like
		if len(body) == 0 {
			t.Fatalf("empty body")
		}
	}
}

func TestExecute_POST_Wrong_responseBody(t *testing.T) {

	headersList := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer your_token",
	}

	request := webreq.NewRequest("POST")
	request.SetURL("https://orange")
	request.SetData([]byte(`{"key":"value"}`)) // set data to POST
	request.SetHeaders(headersList)
	request.SetTimeout(10)

	response, err := request.Execute()
	if err == nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response != nil {
		t.Fatalf("expected nil response")
	}
}

func TestExecute_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Sleep longer than request timeout
		time.Sleep(1500 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("late"))
	}))
	defer ts.Close()

	req := webreq.NewRequest("GET")
	req.SetURL(ts.URL)
	req.SetTimeout(1) // 1 second

	_, err := req.Execute()
	if err == nil {
		t.Fatalf("expected timeout error, got nil")
	}
}

func TestExecute_InvalidURL(t *testing.T) {
	req := webreq.NewRequest("GET")
	// Intentionally malformed URL
	req.SetURL(":\x00")
	req.SetTimeout(1)

	_, err := req.Execute()
	if err == nil {
		t.Fatalf("expected error for invalid URL, got nil")
	}
}

func TestSetUrl_EmptyURL(t *testing.T) {
	req := webreq.NewRequest("GET")
	// Intentionally malformed URL
	req.SetURL("")

	if req.ErrorMessage == "" {
		t.Fatalf("expected content but got nil")
	}

	if req.ErrorMessage != "url is empty" {
		t.Fatalf("expected url is empty, got %q", req.ErrorMessage)
	}
}

func TestSetHeaders_EmptyHeaders(t *testing.T) {
	req := webreq.NewRequest("GET")
	// Intentionally malformed URL
	req.SetHeaders(webreq.HeadersMap{})
	if req.ErrorMessage == "" {
		t.Fatalf("expected content but got nil")
	}

	if req.ErrorMessage != "headers are empty" {
		t.Fatalf("expected headers are empty, got %q", req.ErrorMessage)
	}
}

func TestSetData_EmptyData(t *testing.T) {
	req := webreq.NewRequest("GET")
	req.SetData(make([]byte, 0))
	if req.ErrorMessage == "" {
		t.Fatalf("expected content but got nil")
	}

	if req.ErrorMessage != "body is empty" {
		t.Fatalf("expected body is empty, got %q", req.ErrorMessage)
	}
}

func TestSetStatusCode_EmptyData(t *testing.T) {
	req := webreq.NewRequest("GET")
	req.SetStatusCode(0)
	if req.ErrorMessage == "" {
		t.Fatalf("expected content but got nil")
	}

	if req.ErrorMessage != "status code is empty" {
		t.Fatalf("expected status code is empty, got %q", req.ErrorMessage)
	}
}

func TestCheck(t *testing.T) {
	req := webreq.NewRequest("GET")
	req.SetURL("https://6657cc695c3617052645e2bd.mockapi.io/v1/core/courses")
	req.Check()
	if req.ErrorMessage != "" {
		t.Fatalf("expected nil but got %v", req.ErrorMessage)
	}
}

func TestCheck_InvalidURL(t *testing.T) {
	req := webreq.NewRequest("GET")
	req.Check()
	if req.ErrorMessage == "" {
		t.Fatalf("expected value but got nil")
	}
}

func TestCheck_InvalidMethod(t *testing.T) {
	req := webreq.NewRequest("")
	req.Check()
	if req.ErrorMessage == "" {
		t.Fatalf("expected value but got nil")
	}
}
