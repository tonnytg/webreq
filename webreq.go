package webreq

import (
	"bytes"
	"context"
	"io"
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

func NewHeaders(headers map[string]string) *Headers {
	if len(headers) != 0 {
		return &Headers{
			Headers: headers,
		}
	}
	return &Headers{
		Headers: make(HeadersMap),
	}
}

func (header *Headers) Add(key string, value string) {
	if key == "" || value == "" {
		return
	}
	header.Headers[key] = value
}

type Request struct {
	URL             string
	TimeoutDuration time.Duration
	Headers         HeadersMap
	Method          string
	Data            []byte
	StatusCode      int
	ErrorMessage    string
}

func NewRequest(method string) *Request {

	request := Request{}

	r := request.SetMethod(method)
	if r == nil {
		return nil
	}

	return &Request{
		TimeoutDuration: 10 * time.Second,
		Method:          method,
	}
}

func (request *Request) SetURL(urlValue string) *Request {
	if urlValue == "" {
		request.ErrorMessage = "url is empty"
		return nil
	}
	request.URL = urlValue
	return request
}

func (request *Request) SetTimeout(timeout int) *Request {
	if timeout == 0 {
		request.TimeoutDuration = time.Duration(10) * time.Second
		return request
	}
	request.TimeoutDuration = time.Duration(timeout) * time.Second
	return request
}

func (request *Request) SetHeaders(headers HeadersMap) *Request {
	if len(headers) == 0 {
		request.ErrorMessage = "headers is empty"
		return request
	}
	request.Headers = headers
	return request
}

func (request *Request) SetData(bodyValue []byte) *Request {
	if len(bodyValue) == 0 {
		request.ErrorMessage = "body is empty"
		return request
	}
	request.Data = bodyValue
	return request
}

func (request *Request) SetMethod(requestMethod string) *Request {
	if requestMethod == "" {
		request.ErrorMessage = "request type is empty"
		return nil
	}
	request.Method = requestMethod
	return request
}

func (request *Request) SetStatusCode(statusCodeValue int) *Request {
	if statusCodeValue == 0 {
		request.ErrorMessage = "status code is empty"
		return request
	}
	request.StatusCode = statusCodeValue
	return request
}

func (request *Request) Execute() ([]byte, error) {
	client := &http.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), request.TimeoutDuration)
	defer cancel()

	webRequest, err := http.NewRequestWithContext(ctx, request.Method, request.URL, bytes.NewReader(request.Data))
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
