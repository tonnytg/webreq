package webreq

import (
	"encoding/json"
	"fmt"
)

// only works with this methods
const (
	MethodGet  = "GET"
	MethodPost = "POST"
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

// PrintHeaders print headers parallel with request
func PrintHeaders(headers *Headers) {
	for _, header := range headers.List {
		for k, v := range header {
			fmt.Printf("%s: %s\n", k, v)
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
