package webreq_test

import (
	"github.com/tonnytg/webreq"
	"testing"
)

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

func TestWrongCall(t *testing.T) {

	request := webreq.NewRequest("GET")

	body, err := request.Execute()
	if err == nil {
		t.Error("Expected an error but didn't get one")
	}
	bodyString := string(body)
	if bodyString != "" {
		t.Error("body is not empty")
	}
}
