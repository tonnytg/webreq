package webreq

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

// Get receive an url, you can send headers and timeout parameters for request.
func Get(url string, headers *Headers, timeOut int) ([]byte, error) {

	client := &http.Client{
		CheckRedirect: nil,
	}
	if timeOut == 0 {
		timeOut = 10
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeOut)*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	go PrintHeaders(headers)
	for _, headers := range headers.List {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("error to close body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
