package webreq_test

import (
	"fmt"
	"github.com/tonnytg/webreq"
	"testing"
	"time"
)

func TestPackageCall(t *testing.T) {

	var headersList = make(map[string]string)
	headersList["Content-Type"] = "application/json"

	headers := webreq.NewHeaders(headersList)

	request := webreq.NewRequest("GET")
	request.SetURL("https://examples.com/values")
	request.SetHeaders(headers.ListHeaders) // Pass the map directly here
	request.SetTimeout(10)

	body, err := request.Execute()
	if err != nil {
		t.Error(err)
	}
	bodyString := string(body)
	if bodyString == "" {
		t.Error("body is empty")
	}
}

// TestSetURL verify URL it is correctly defined
func TestSetURL(t *testing.T) {
	t.Run("Non-empty URL", func(t *testing.T) {
		request := webreq.NewRequest("GET")
		request.SetURL("https://example.com")

		// Verifique se a URL é definida corretamente
		if request.URL != "https://example.com" {
			t.Errorf("URL not set correctly")
		}
	})
}

func TestSetTimeout(t *testing.T) {
	request := webreq.NewRequest("GET")
	request.SetTimeout(10)
	if request.TimeoutDuration != (10 * time.Second) {
		fmt.Println(request.TimeoutDuration)
		t.Error("request.TimeoutDuration is not 10 * time.Second")
	}
	return
}

func TestSetHeaders(t *testing.T) {

	var headersList = make(map[string]string)
	headersList["Content-Type"] = "application/json"

	headers := webreq.NewHeaders(headersList)
	if headers.ListHeaders["Content-Type"] != "application/json" {
		t.Error("headers.Headers[Content-Type] is not application/json")
	}

	request := webreq.NewRequest("GET")
	request.SetHeaders(headers.ListHeaders)
	if request.ErrorMessage != "" {
		t.Error("request.ErrorMessage is not empty")
	}
	return
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
	return
}

func TestSetData(t *testing.T) {

	request := webreq.NewRequest("GET")
	request.SetData([]byte("test"))
	if request.ErrorMessage != "" {
		t.Error("request.ErrorMessage is not empty")
	}
	return
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
	return
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
	return
}

func TestSetNewRequest(t *testing.T) {

	newRequest := webreq.NewRequest("GET")
	if newRequest == nil {
		t.Error("newRequest is nil")
	}
}

func TestEndToEnd(t *testing.T) {

	headers := webreq.NewHeaders(nil)
	headers.Add("Content-Type", "application/json")

	request := webreq.NewRequest("GET")
	if request == nil {
		t.Error("request is nil")
	}
	request.SetURL("https://example.com/values")
	if request.ErrorMessage != "" {
		t.Error("request.ErrorMessage is not empty")
	}
	request.SetHeaders(headers.ListHeaders)
	if request.ErrorMessage != "" {
		t.Error("request.ErrorMessage is not empty")
	}
	request.SetTimeout(10)

	body, err := request.Execute()
	if err != nil {
		t.Error(err)
	}
	bodyString := string(body)
	if bodyString == "" {
		t.Error("body is empty")
	}
}
