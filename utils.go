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

// StructToBytes convert a struct to bytes and return error if it fails
func StructToBytes(data interface{}) ([]byte, error) {

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}
