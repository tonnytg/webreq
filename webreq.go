package webreq

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// only works with this methods
const (
	Method = "GET"
)

// Get receive an url, you can send headers and timeout parameters for request.
func Get(url string, headers map[string]string, timeOut int) ([]byte, error) {

	fmt.Println("url:", url)
	// print headers
	fmt.Println("headers:")
	if len(headers) > 0 {
		for k, v := range headers {
			fmt.Println("\t", k, ":", v)
		}
	} else {
		fmt.Println("\tNo headers")
	}

	// client run everything
	client := &http.Client{
		CheckRedirect: nil,
	}

	// context help us to control timeout for request
	ctx := context.Background()
	if timeOut == 0 {
		timeOut = 10
	}
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeOut)*time.Second)
	defer cancel()

	// request with NewRequest permit add headers
	request, err := http.NewRequestWithContext(ctx, Method, url, nil)
	if err != nil {
		return nil, err
	}

	// loop to add each header in request
	for _, v := range headers {
		request.Header.Add("If-None-Match", v)
	}

	// execute call and return *http.Response type
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	// clone request after process
	defer response.Body.Close()

	// convert information to slice of byte
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
