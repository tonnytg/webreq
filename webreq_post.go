package webreq

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

// Post you can send a struct and receive a response from a url
func Post(url string, data []byte, headers map[string]string, timeOut int) ([]byte, error) {

	client := &http.Client{
		CheckRedirect: nil,
	}
	if timeOut == 0 {
		timeOut = 10
	}
	dataBody := io.NopCloser(bytes.NewReader(data))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeOut)*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, MethodPost, url, dataBody)
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
