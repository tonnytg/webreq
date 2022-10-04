package webreq_test

import (
	"github.com/tonnytg/webreq"
	"testing"
)

func TestPackageGet(t *testing.T) {
	url := "https://www.google.com/robots.txt"
	timeOut := 20
	headers := []string{}
	_, err := webreq.Get(url, headers, timeOut)
	if err != nil {
		t.Error(err)
	}
}
