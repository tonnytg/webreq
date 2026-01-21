package webreq

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	MethodGet  = "GET"
	MethodPost = "POST"
	// DefaultMaxResponseSize is the default maximum size for response bodies (100MB)
	DefaultMaxResponseSize = 100 * 1024 * 1024 // 100MB
)

var (
	// defaultClient is a shared HTTP client with connection pooling for better performance
	defaultClient     *http.Client
	defaultClientOnce sync.Once
)

// getDefaultClient returns a shared HTTP client with optimized transport settings
func getDefaultClient() *http.Client {
	defaultClientOnce.Do(func() {
		transport := &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		}
		defaultClient = &http.Client{
			Transport: transport,
		}
	})
	return defaultClient
}

type HeadersMap map[string]string

type Headers struct {
	ListHeaders HeadersMap
}

// NewHeaders creates a new Headers instance, initializing with provided headers or an empty map
func NewHeaders(headers map[string]string) *Headers {
	if headers != nil {
		return &Headers{
			ListHeaders: headers,
		}
	}
	return &Headers{
		ListHeaders: make(HeadersMap),
	}
}

// Add adds a new header key-value pair to the Headers
func (header *Headers) Add(key string, value string) {
	if key != "" && value != "" {
		header.ListHeaders[key] = value
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
	MaxResponseSize int64 // Maximum size for response body in bytes
}

// NewRequest creates a new Request with the specified method
func NewRequest(method string) *Request {
	return &Request{
		TimeoutDuration: 10 * time.Second,
		Method:          method,
		MaxResponseSize: DefaultMaxResponseSize,
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

// SetMaxResponseSize sets the maximum response body size in bytes
func (request *Request) SetMaxResponseSize(size int64) *Request {
	if size > 0 {
		request.MaxResponseSize = size
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

func (request *Request) Check() error {

	if request.URL == "" {
		request.ErrorMessage = "url is empty"
	}

	if request.Method == "" {
		request.ErrorMessage = "method is empty"
	}

	return nil
}

// Execute sends the request and returns the response body and error if any
func (request *Request) Execute() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), request.TimeoutDuration)
	defer cancel()
	return request.ExecuteWithContext(ctx)
}

// ExecuteWithContext sends the request with a custom context and returns the response body and error if any
func (request *Request) ExecuteWithContext(ctx context.Context) ([]byte, error) {
	client := getDefaultClient()

	var body io.Reader
	if len(request.Data) > 0 {
		body = bytes.NewReader(request.Data)
	}

	webRequest, err := http.NewRequestWithContext(ctx, request.Method, request.URL, body)
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

	// Limit response body size to prevent memory exhaustion attacks
	limitedReader := io.LimitReader(response.Body, request.MaxResponseSize)
	responseBody, err := io.ReadAll(limitedReader)
	if err != nil {
		return nil, err
	}

	request.StatusCode = response.StatusCode
	return responseBody, nil
}
