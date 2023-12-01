package webreq_test

import (
	"github.com/tonnytg/webreq"
	"testing"
)

func TestPackageCall(t *testing.T) {

	var headersList = make(map[string]string)
	headersList["Content-Type"] = "application/json"

	headers := webreq.NewHeaders(headersList)

	request := webreq.NewRequest("GET")
	request.SetURL("https://610aa52552d56400176afebe.mockapi.io/api/v1/friendlist")
	request.SetHeaders(headers.Headers) // Pass the map directly here
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

func TestSetURL(t *testing.T) {
	// Teste para verificar se a URL é definida corretamente quando não está vazia
	t.Run("Non-empty URL", func(t *testing.T) {
		request := webreq.NewRequest("GET")
		request.SetURL("https://example.com")

		// Verifique se a URL é definida corretamente
		if request.URL != "https://example.com" {
			t.Errorf("URL not set correctly")
		}
	})
}

func TestSetNewRequest(t *testing.T) {

	newRequest := webreq.NewRequest("GET")
	if newRequest == nil {
		t.Error("newRequest is nil")
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
	return
}

func TestEndToEnd(t *testing.T) {

	headers := webreq.NewHeaders(nil)
	headers.Add("Content-Type", "application/json")

	request := webreq.NewRequest("GET")
	if request == nil {
		t.Error("request is nil")
	}
	request.SetURL("https://610aa52552d56400176afebe.mockapi.io/api/v1/friendlist")
	if request.ErrorMessage != "" {
		t.Error("request.ErrorMessage is not empty")
	}
	request.SetHeaders(headers.Headers)
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
