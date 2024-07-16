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

// NewHeaders creates a new Headers instance, initializing with provided headers or an empty map
func NewHeaders(headers map[string]string) *Headers {
	if headers != nil {
		return &Headers{
			Headers: headers,
		}
	}
	return &Headers{
		Headers: make(HeadersMap),
	}
}

// Add adds a new header key-value pair to the Headers
func (header *Headers) Add(key string, value string) {
	if key != "" && value != "" {
		header.Headers[key] = value
	}
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

// NewRequest creates a new Request with the specified method
func NewRequest(method string) *Request {
	return &Request{
		TimeoutDuration: 10 * time.Second,
		Method:          method,
	}
}

// SetURL sets the URL of the request
func (request *Request) SetURL(urlValue string) *Request {
	if urlValue == "" {
		request.ErrorMessage = "url is empty"
		return nil
	}
	request.URL = urlValue
	return request
}

// SetTimeout sets the request timeout duration
func (request *Request) SetTimeout(timeout int) *Request {
	if timeout > 0 {
		request.TimeoutDuration = time.Duration(timeout) * time.Second
	}
	return request
}

// SetHeaders sets the headers of the request
func (request *Request) SetHeaders(headers HeadersMap) *Request {
	if len(headers) > 0 {
		request.Headers = headers
	} else {
		request.ErrorMessage = "headers are empty"
	}
	return request
}

// SetData sets the data to be sent with the request
func (request *Request) SetData(bodyValue []byte) *Request {
	if len(bodyValue) > 0 {
		request.Data = bodyValue
	} else {
		request.ErrorMessage = "body is empty"
	}
	return request
}

// SetMethod sets the method of the request
func (request *Request) SetMethod(requestMethod string) *Request {
	if requestMethod != "" {
		request.Method = requestMethod
	} else {
		request.ErrorMessage = "request method is empty"
	}
	return request
}

// SetStatusCode sets the status code of the response
func (request *Request) SetStatusCode(statusCodeValue int) *Request {
	if statusCodeValue > 0 {
		request.StatusCode = statusCodeValue
	} else {
		request.ErrorMessage = "status code is empty"
	}
	return request
}

// Execute sends the request and returns the response body and error if any
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

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	request.StatusCode = response.StatusCode
	return responseBody, nil
}
