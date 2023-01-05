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

// PrintHeaders print headers paralelly with request
func PrintHeaders(headers map[string]string) {
	fmt.Println("headers:")
	if len(headers) > 0 {
		for k, v := range headers {
			fmt.Println("\t", k, ":", v)
		}
	} else {
		fmt.Println("\tNo headers")
	}
}

// Get receive an url, you can send headers and timeout parameters for request.
func Get(url string, headers map[string]string, timeOut int) ([]byte, error) {

	client := &http.Client{
		CheckRedirect: nil,
	}
	if timeOut == 0 {
		timeOut = 10
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeOut)*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, Method, url, nil)
	if err != nil {
		return nil, err
	}
	go PrintHeaders(headers)
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
