package webreq_test

import (
	"github.com/tonnytg/webreq"
	"testing"
)

// fakeWriter é uma implementação simples de io.Writer para capturar a saída do log
type fakeWriter struct {
	target *string
}

func (fw *fakeWriter) Write(p []byte) (n int, err error) {
	*fw.target = string(p)
	return len(p), nil
}

func TestPackageCall(t *testing.T) {

	headers := webreq.NewHeaders()
	headers.Add("Content-Type", "application/json")

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
