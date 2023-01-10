package webreq_test

import (
	"github.com/tonnytg/webreq"
	"testing"
)

func TestPackageGet(t *testing.T) {
	url := "https://www.google.com"
	timeOut := 20

	headers := webreq.NewHeaders()
	headers.Add("Content-Type", "application/json")

	resp, err := webreq.Execute("GET", url, headers, nil, timeOut)
	if err != nil {
		t.Error(err)
	}
	body := string(resp)
	if body == "" {
		t.Error("body is empty")
	}
}
