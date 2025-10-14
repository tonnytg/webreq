package webreq_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/tonnytg/webreq"
)

func TestPackageCall(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/values" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Fatalf("unexpected content type: %s", r.Header.Get("Content-Type"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	headersList := map[string]string{"Content-Type": "application/json"}
	headers := webreq.NewHeaders(headersList)

	request := webreq.NewRequest("GET")
	request.SetURL(ts.URL + "/values")
	request.SetHeaders(headers.ListHeaders)
	request.SetTimeout(2)

	body, err := request.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) == "" {
		t.Fatal("body is empty")
	}
}

// TestSetURL verify URL it is correctly defined
func TestSetURL(t *testing.T) {
	t.Run("Non-empty URL", func(t *testing.T) {
		request := webreq.NewRequest("GET")
		request.SetURL("https://example.com")

		if request.URL != "https://example.com" {
			t.Errorf("URL not set correctly")
		}
	})
}

func TestSetTimeout(t *testing.T) {
	request := webreq.NewRequest("GET")
	request.SetTimeout(10)
	if request.TimeoutDuration != 10*time.Second {
		fmt.Println(request.TimeoutDuration)
		t.Error("request.TimeoutDuration is not 10 * time.Second")
	}
}

func TestSetHeaders(t *testing.T) {
	headersList := map[string]string{"Content-Type": "application/json"}

	headers := webreq.NewHeaders(headersList)
	if headers.ListHeaders["Content-Type"] != "application/json" {
		t.Error("headers.Headers[Content-Type] is not application/json")
	}

	request := webreq.NewRequest("GET")
	request.SetHeaders(headers.ListHeaders)
	if request.ErrorMessage != "" {
		t.Error("request.ErrorMessage is not empty")
	}
}

func TestSetHeaders2(t *testing.T) {
	headers := webreq.NewHeaders(nil)
	headers.Add("Content-Type", "application/json")
	if headers.ListHeaders["Content-Type"] != "application/json" {
		t.Error("headers.Headers[Content-Type] is not application/json")
	}

	request := webreq.NewRequest("GET")
	request.SetHeaders(headers.ListHeaders)
	if request.ErrorMessage != "" {
		t.Error("request.ErrorMessage is not empty")
	}
}

func TestSetData(t *testing.T) {
	request := webreq.NewRequest("GET")
	request.SetData([]byte("test"))
	if request.ErrorMessage != "" {
		t.Error("request.ErrorMessage is not empty")
	}
}

func TestSetMethod(t *testing.T) {
	request := webreq.NewRequest("GET")
	if request == nil {
		t.Error("request is nil")
	}
	request.SetMethod("GET")
	if request.Method != "GET" {
		t.Error("request.Method is not GET")
	}

	request.SetMethod("POST")
	if request.Method != "POST" {
		t.Error("request.Method is not POST")
	}

	request.SetMethod("")
	if request.ErrorMessage != "request method is empty" {
		t.Error("request.Method is not empty")
	}
}

func TestSetStatusCode(t *testing.T) {
	request := webreq.NewRequest("GET")
	if request == nil {
		t.Error("request is nil")
	}
	request.SetStatusCode(200)
	if request.StatusCode != 200 {
		t.Error("request.StatusCode is not 200")
	}
}

func TestSetNewRequest(t *testing.T) {
	newRequest := webreq.NewRequest("GET")
	if newRequest == nil {
		t.Error("newRequest is nil")
	}
}

func TestEndToEnd(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Fatalf("unexpected content type: %s", r.Header.Get("Content-Type"))
		}
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"status": "ok"}
		_ = json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	headers := webreq.NewHeaders(nil)
	headers.Add("Content-Type", "application/json")

	request := webreq.NewRequest("GET")
	if request == nil {
		t.Error("request is nil")
	}
	request.SetURL(ts.URL)
	if request.ErrorMessage != "" {
		t.Error("request.ErrorMessage is not empty")
	}
	request.SetHeaders(headers.ListHeaders)
	if request.ErrorMessage != "" {
		t.Error("request.ErrorMessage is not empty")
	}
	request.SetTimeout(2)

	body, err := request.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(body) == 0 {
		t.Fatal("body is empty")
	}
}
