package main

import (
	"encoding/json"
	"fmt"
	"github.com/tonnytg/webreq"
	"time"
)

// Friend represents a friend with a creation date and name
type Friend struct {
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func main() {

	var headers = make(map[string]string)
	headers["Content-Type"] = "application/json"

	// Create a new Friend instance
	data := Friend{
		CreatedAt: time.Now(),
		Name:      "Tonny",
	}

	// Convert Friend instance to JSON bytes
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Create a new POST request
	request := webreq.NewRequest("POST")
	// Set the request URL
	request.SetURL("https://example.com/api/v1/values")
	// Set the request body data
	request.SetData(dataBytes)
	// Set the request headers
	request.SetHeaders(headers) // Set map directly
	// Set the request timeout to 10 seconds
	request.SetTimeout(10)

	// Execute the request and get the response
	response, err := request.Execute()
	if err != nil {
		fmt.Println("Error executing request:", err)
		return
	}

	// Convert response bytes to string
	bodyString := string(response)
	if bodyString == "" {
		fmt.Println("Response status code:", request.StatusCode)
		fmt.Println("Response body:", bodyString)
	} else {
		fmt.Println("Response status code:", request.StatusCode)
		fmt.Println("Response body:", bodyString)
	}
}
