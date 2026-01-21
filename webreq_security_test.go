package webreq

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestMaxResponseSize_Default tests that the default max response size is set
func TestMaxResponseSize_Default(t *testing.T) {
	request := NewRequest(MethodGet)
	if request.MaxResponseSize != DefaultMaxResponseSize {
		t.Errorf("Expected MaxResponseSize to be %d, got %d", DefaultMaxResponseSize, request.MaxResponseSize)
	}
}

// TestSetMaxResponseSize tests setting a custom max response size
func TestSetMaxResponseSize(t *testing.T) {
	request := NewRequest(MethodGet)
	customSize := int64(1024 * 1024) // 1MB
	request.SetMaxResponseSize(customSize)
	
	if request.MaxResponseSize != customSize {
		t.Errorf("Expected MaxResponseSize to be %d, got %d", customSize, request.MaxResponseSize)
	}
}

// TestSetMaxResponseSize_Zero tests that zero size is ignored
func TestSetMaxResponseSize_Zero(t *testing.T) {
	request := NewRequest(MethodGet)
	originalSize := request.MaxResponseSize
	request.SetMaxResponseSize(0)
	
	if request.MaxResponseSize != originalSize {
		t.Errorf("Expected MaxResponseSize to remain %d, got %d", originalSize, request.MaxResponseSize)
	}
}

// TestSetMaxResponseSize_Negative tests that negative size is ignored
func TestSetMaxResponseSize_Negative(t *testing.T) {
	request := NewRequest(MethodGet)
	originalSize := request.MaxResponseSize
	request.SetMaxResponseSize(-1000)
	
	if request.MaxResponseSize != originalSize {
		t.Errorf("Expected MaxResponseSize to remain %d, got %d", originalSize, request.MaxResponseSize)
	}
}

// TestMaxResponseSize_LimitEnforced tests that response body is limited
func TestMaxResponseSize_LimitEnforced(t *testing.T) {
	// Create a test server that returns a large response
	largeData := make([]byte, 2000) // 2KB of data
	for i := range largeData {
		largeData[i] = 'A'
	}
	
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(largeData)
	}))
	defer server.Close()
	
	// Set max response size to 1KB
	request := NewRequest(MethodGet)
	request.SetURL(server.URL)
	request.SetMaxResponseSize(1000)
	
	response, err := request.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
	
	// Should only receive up to 1000 bytes
	if len(response) != 1000 {
		t.Errorf("Expected response size to be limited to 1000 bytes, got %d", len(response))
	}
}

// TestMaxResponseSize_SmallResponse tests that small responses work correctly
func TestMaxResponseSize_SmallResponse(t *testing.T) {
	smallData := []byte("Hello, World!")
	
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(smallData)
	}))
	defer server.Close()
	
	request := NewRequest(MethodGet)
	request.SetURL(server.URL)
	request.SetMaxResponseSize(1000)
	
	response, err := request.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
	
	// Should receive the full response
	if len(response) != len(smallData) {
		t.Errorf("Expected response size to be %d, got %d", len(smallData), len(response))
	}
	
	if string(response) != string(smallData) {
		t.Errorf("Expected response to be %q, got %q", smallData, response)
	}
}
