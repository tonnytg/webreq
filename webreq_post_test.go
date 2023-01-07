package webreq_test

import (
	"fmt"
	"github.com/tonnytg/webreq"
	"testing"
	"time"
)

type Friend struct {
	CreatedAt  time.Time `json:"createdAt"`
	Name       string    `json:"name"`
	Job        string    `json:"job"`
	FamilyName string    `json:"familyName"`
}

func TestPackagePost(t *testing.T) {
	url := "https://610aa52552d56400176afebe.mockapi.io/api/v1/friendlist"
	timeOut := 20

	headers := webreq.NewHeaders()
	headers.Add("Content-Type", "application/json")

	f := Friend{
		CreatedAt:  time.Now(),
		Name:       "Tonny",
		Job:        "Developer",
		FamilyName: "Gomes",
	}
	// convert f to bytes
	fBytes, err := webreq.StructToBytes(f)
	if err != nil {
		t.Error(err)
	}
	body, err := webreq.Post(url, headers, timeOut, fBytes)
	if err != nil {
		t.Error(err)
	}
	bodyString := string(body)
	if bodyString == "" {
		t.Error("body is empty")
	}
	fmt.Println(bodyString)
}
