package pkggowebreq_test

import (
	"github.com/tonnytg/pkggowebreq"
	"testing"
)

func TestPackageGet(t *testing.T) {
	url := "https://610aa52552d56400176afebe.mockapi.io/api/v1/friendlist"
	timeOut := 20
	headers := []string{}
	_, err := pkggowebreq.Get(url, headers, timeOut)
	if err != nil {
		t.Error(err)
	}
}
