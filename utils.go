package webreq

import (
	"encoding/json"
)

// only works with this methods
const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

type H map[string]string

type Headers struct {
	Headers []H
}

func (h *Headers) Add(header H) {
	h.Headers = append(h.Headers, header)
}

// PrintHeaders print headers paralelly with request
func PrintHeaders(headers Headers) {
	for _, header := range headers.Headers {
		for k, v := range header {
			println(k, v)
		}
	}
}

// StructToBytes convert a struct to bytes and return error if it fails
func StructToBytes(data interface{}) ([]byte, error) {

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}
