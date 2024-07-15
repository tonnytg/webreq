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

	headers := webreq.NewHeaders(nil)
	headers.Add("Content-Type", "application/json")

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
	request.SetData(fBytes)
	request.SetHeaders(headers.Headers) // Set map directly
	request.SetTimeout(10)

	response, err := request.Execute()
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(response)
	if bodyString == "" {
		fmt.Println("response status code:", request.StatusCode)
		fmt.Println("response body:", bodyString)
	}

	fmt.Println("response status code:", request.StatusCode)
	fmt.Println("response body:", bodyString)
}
