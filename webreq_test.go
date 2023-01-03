package webreq_test

import (
	"github.com/tonnytg/webreq"
	"testing"
)

func TestPackageGet(t *testing.T) {
	url := "https://www.google.com"
	timeOut := 20
	headers := map[string]string{}
	body, err := webreq.Get(url, headers, timeOut)
	if err != nil {
		t.Error(err)
	}
	bodyString := string(body)
	if bodyString == "" {
		t.Error("body is empty")
	}
	//fmt.Println(bodyString)
}
