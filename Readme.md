# Web Request for Go Packages
[![Go Report Card](https://goreportcard.com/badge/github.com/tonnytg/webreq)](https://goreportcard.com/report/github.com/tonnytg/webreq) [![codecov](https://codecov.io/gh/tonnytg/webreq/branch/main/graph/badge.svg?token=PYI6QQKGTV)](https://codecov.io/gh/tonnytg/webreq) ![example workflow](https://github.com/tonnytg/webreq/actions/workflows/go.yml/badge.svg) 

## What problem solve?
Create easily web request for APIs with Headers using methods `GET` or `POST`

## What is this?

This is module to help you make web requests in Go, it is a wrapper around the standard library's `http` package.
You can use Get or Post to make a request, and then use the `Response` object to get the response body, headers, status code, etc.

### Example

    package main
    
    import (
        "log"
        "github.com/tonnytg/webreq"
    )

    func main() {
    
        request := webreq.NewRequest("GET")
        request.SetURL("https://6657cc695c3617052645e2bd.mockapi.io/v1/core/courses")
    
        data, err := request.Execute()
        if err != nil {
            log.Fatalf("Request failed: %v", err)
            return
        }
        
        log.Println(string(data))
    }


## Why

So many times I needed to make a request in an API and after convert body to struct, and every time I needed to configure http.Client and put headers and body. Thinking about "don't repeat your self" I believe this unified code can gain time build an easy way to make this request.

What do you think? Liked, set star in this project to help me to help others!

### Using only Http Get

    client := &http.Client{
        Timeout: 10 * time.Second,
    }
    
    req, err := http.NewRequest("GET", "https://api.example.com/data", nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }
    req.Header.Add("Authorization", "Bearer your_token")
    
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error executing request:", err)
        return
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }
    fmt.Println("Response:", string(body))

### Using WebReq

	var headers = make(map[string]string)
	headers["Content-Type"] = "application/json"
    headers["Authorization"] = "Bearer your_token"
    
    request := webreq.NewRequest("POST")
    request.SetURL("https://api.example.com/data")
    request.SetData([]byte(`{"key":"value"}`))
    request.SetHeaders(headers)
    request.SetTimeout(10)
    
    response, err := request.Execute()
    if err != nil {
        fmt.Println("Error executing request:", err)
        return
    }
    fmt.Println("Response:", string(response))

## Install

> go get github.com/tonnytg/webreq

## Basic Usage

### Get

Create a slice of headers to use and input in request 

    package main
    
    import (
        "log"
        "github.com/tonnytg/webreq"
    )
    
    func main() {
        request := webreq.NewRequest("GET")
        request.SetURL("https://example.com")
    
        // Execute the request and get the response
        response, err := request.Execute()
        if err != nil {
            fmt.Println("Error executing request:", err)
            return
        }

        log.Println(string(response))
    }


### Post

Create a body data to send in request

    package main
    
    import (
        "log"
        "github.com/tonnytg/webreq"
    )
    
    func main() {
        var headers = make(map[string]string)
        headers["Content-Type"] = "application/json"
        headers["Authorization"] = "Bearer your_token"
        
        request := webreq.NewRequest("POST")
        request.SetURL("https://api.example.com/data")
        request.SetData([]byte(`{"key":"value"}`)) // set data to POST
        request.SetHeaders(headers)
        request.SetTimeout(10)
        
        response, err := request.Execute()
        if err != nil {
            log.Println("Error executing request:", err)
            return
        }
        log.Println("Response:", string(response))
    }
        


## Advange Usage


### Custom Headers

	headersList := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer your_token",
	}
	
	request := webreq.NewRequest("POST")
	request.SetURL("https://6657cc695c3617052645e2bd.mockapi.io/v1/core/courses")
	request.SetData([]byte(`{"key":"value"}`)) // set data to POST
	request.SetHeaders(headersList)
	request.SetTimeout(10)
	
	response, err := request.Execute()
	if err != nil {
		log.Println("Error executing request:", err)
		return
	}
	log.Println("Response:", string(response))

### Custom Context (for cancellation, tracing, etc.)

	package main
	
	import (
		"context"
		"log"
		"time"
		"github.com/tonnytg/webreq"
	)
	
	func main() {
		// Create a context with custom timeout or cancellation
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		
		request := webreq.NewRequest("GET")
		request.SetURL("https://api.example.com/data")
		
		// Use ExecuteWithContext for custom context control
		response, err := request.ExecuteWithContext(ctx)
		if err != nil {
			log.Println("Error executing request:", err)
			return
		}
		log.Println("Response:", string(response))
	}

## Performance

WebReq is optimized for performance with:
- **Connection Pooling**: Shared HTTP client with connection reuse (up to 100 idle connections)
- **HTTP Keep-Alive**: Reduces latency for multiple requests to the same host
- **Efficient Memory Usage**: ~6KB allocation per request with optimized body handling
- **Concurrent Performance**: 5-6x faster under parallel load thanks to connection pooling

See [PERFORMANCE_IMPROVEMENTS.md](PERFORMANCE_IMPROVEMENTS.md) for detailed benchmarks and optimization details.

## Contributing
- Fork the repository
- Create a new branch (git checkout -b feature-foo)
- Commit your changes (git commit -am 'Add some foo')
- Push to the branch (git push origin feature-foo)
- Create a new Pull Request


## License
This project is licensed under the MIT License.


## Contact

For any inquiries, please open an issue or reach out to the maintainer.

email: tonnytg@gmail.com
