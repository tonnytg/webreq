package webreq

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

type H map[string]string

type Headers struct {
	List []H
}

func NewHeaders() *Headers {
	return &Headers{}
}

func (h *Headers) Add(key string, value string) {
	h.List = append(h.List, H{key: value})
}

type Request struct {
	URL         string
	TimeOut     int
	Headers     []H
	TypeRequest string
	Body        []byte
	StatusCode  int
}

func Builder(method string) *Request {
	r := Request{}

	r.SetTimeOut(10)
	r.SetTypeRequest(method)

	return &r
}

func (r *Request) SetURL(url string) *Request {
	r.URL = url
	return r
}

func (r *Request) SetTimeOut(timeOut int) *Request {
	r.TimeOut = timeOut
	return r
}

func (r *Request) SetHeaders(headers *Headers) *Request {
	r.Headers = headers.List
	return r
}

func (r *Request) SetBody(body []byte) *Request {
	r.Body = body
	return r
}

func (r *Request) SetTypeRequest(typeRequest string) *Request {
	r.TypeRequest = typeRequest
	return r
}

func (r *Request) SetStatusCode(statusCode int) *Request {
	r.StatusCode = statusCode
	return r
}

func (r *Request) Execute() ([]byte, error) {

	client := &http.Client{}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(r.TimeOut)*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx,
		r.TypeRequest,
		r.URL,
		bytes.NewReader(r.Body))
	if err != nil {
		return nil, err
	}
	for _, header := range r.Headers {
		for key, value := range header {
			request.Header.Add(key, value)
		}
	}

	resp, err := client.Do(request)
	r.SetStatusCode(resp.StatusCode)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println("error to close body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
