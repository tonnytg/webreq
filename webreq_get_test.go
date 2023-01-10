package webreq_test

import (
	"github.com/tonnytg/webreq"
	"testing"
)

func TestPackageCall(t *testing.T) {

	headers := webreq.NewHeaders()
	headers.Add("Content-Type", "application/json")

	request := webreq.Builder("GET")
	request.SetURL("https://610aa52552d56400176afebe.mockapi.io/api/v1/friendlist")

	_, err := request.Execute()
	if err != nil {
		t.Error(err)
	}
}
