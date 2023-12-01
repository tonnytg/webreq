package webreq_test

import (
	"encoding/json"
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

	headers := webreq.NewHeaders(nil)
	headers.Add("Content-Type", "application/json")
	if len(headers.Headers) != 1 {
		t.Error("headers is empty")
	}

	f := Friend{
		CreatedAt:  time.Now(),
		Name:       "Tonny",
		Job:        "Developer",
		FamilyName: "Gomes",
	}

	// convert f to bytes
	fBytes, err := json.Marshal(f)
	if err != nil {
		t.Error(err)
	}

	request := webreq.NewRequest("POST")
	request.SetURL("https://610aa52552d56400176afebe.mockapi.io/api/v1/friendlist")
	request.SetData(fBytes)
	request.SetHeaders(headers.Headers) // Set map directly
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
