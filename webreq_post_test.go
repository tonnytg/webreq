package webreq_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/tonnytg/webreq"
)

type Friend struct {
	CreatedAt  time.Time `json:"createdAt"`
	Name       string    `json:"name"`
	Job        string    `json:"job"`
	FamilyName string    `json:"familyName"`
}

func TestPackagePost(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/values" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Fatalf("unexpected content type: %s", ct)
		}
		defer r.Body.Close()
		var payload Friend
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("failed to decode body: %v", err)
		}
		if payload.Name == "" {
			t.Fatal("expected name in payload")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]any{"received": payload.Name})
	}))
	defer ts.Close()

	headers := webreq.NewHeaders(nil)
	headers.Add("Content-Type", "application/json")
	if len(headers.ListHeaders) != 1 {
		t.Error("headers is empty")
	}

	f := Friend{
		CreatedAt:  time.Now(),
		Name:       "Tonny",
		Job:        "Developer",
		FamilyName: "Gomes",
	}

	fBytes, err := json.Marshal(f)
	if err != nil {
		t.Fatal(err)
	}

	request := webreq.NewRequest("POST")
	request.SetURL(ts.URL + "/values")
	request.SetData(fBytes)
	request.SetHeaders(headers.ListHeaders)
	request.SetTimeout(2)

	body, err := request.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(body) == 0 {
		t.Fatal("body is empty")
	}
}
