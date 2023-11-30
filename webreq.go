package webreq

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

type HeadersMap map[string]string

type Headers struct {
	Headers HeadersMap
}

func NewHeaders() *Headers {
	return &Headers{
		Headers: make(HeadersMap),
	}
}

func (header *Headers) Add(key string, value string) {
	header.Headers[key] = value
}

type Request struct {
	URL             string
	TimeoutDuration time.Duration
	Headers         HeadersMap
	RequestType     string
	RequestBody     []byte
	StatusCode      int
}

func NewRequest(method string) Request {
	return Request{
		TimeoutDuration: 10 * time.Second,
		RequestType:     method,
	}
}

func (request *Request) SetURL(urlValue string) *Request {
	if urlValue == "" {
		log.Println("URL cannot be empty")
		return nil
	}
	request.URL = urlValue
	return request
}

func (request *Request) SetTimeout(timeout int) *Request {
	request.TimeoutDuration = time.Duration(timeout) * time.Second
	return request
}

func (request *Request) SetHeaders(headers HeadersMap) *Request {
	request.Headers = headers
	return request
}

func (request *Request) SetBody(bodyValue []byte) *Request {
	request.RequestBody = bodyValue
	return request
}

func (request *Request) SetRequestType(requestTypeValue string) *Request {
	request.RequestType = requestTypeValue
	return request
}

func (request *Request) SetStatusCode(statusCodeValue int) *Request {
	request.StatusCode = statusCodeValue
	return request
}

func (request *Request) Execute() ([]byte, error) {
	client := &http.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), request.TimeoutDuration)
	defer cancel()

	webRequest, err := http.NewRequestWithContext(ctx, request.RequestType, request.URL, bytes.NewReader(request.RequestBody))
	if err != nil {
		return nil, err
	}

	for key, value := range request.Headers {
		webRequest.Header.Add(key, value)
	}

	response, err := client.Do(webRequest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response != nil {
		request.SetStatusCode(response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
