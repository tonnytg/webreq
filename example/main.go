package main

import (
	"encoding/json"
	"fmt"
	"github.com/tonnytg/webreq"
	"time"
)

type Friend struct {
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func main() {

	headers := webreq.NewHeaders()
	headers.Add("Content-Type", "application/json")
	if len(headers.Headers) != 1 {
		fmt.Println("headers is empty")
	}

	f := Friend{
		CreatedAt: time.Now(),
		Name:      "Tonny",
	}

	// convert f to bytes
	fBytes, err := json.Marshal(f)
	if err != nil {
		fmt.Println(err)
	}

	request := webreq.NewRequest("POST")
	request.SetURL("https://623a666d5f037c136217238f.mockapi.io/api/v1/categories")
	request.SetBody(fBytes)
	request.SetHeaders(headers.Headers) // Set map directly
	request.SetTimeout(10)

	body, err := request.Execute()
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(body)
	if bodyString == "" {
		fmt.Println("body is empty")
	}

	fmt.Println(bodyString)
}
