# Web Request for Go Packages
[![Go Report Card](https://goreportcard.com/badge/github.com/tonnytg/webreq)](https://goreportcard.com/report/github.com/tonnytg/webreq) [![codecov](https://codecov.io/gh/tonnytg/webreq/branch/main/graph/badge.svg?token=PYI6QQKGTV)](https://codecov.io/gh/tonnytg/webreq) ![example workflow](https://github.com/tonnytg/webreq/actions/workflows/go.yml/badge.svg) 

## What is this?

This is module to help you make web requests in Go, it is a wrapper around the standard library's `http` package.
You can use Get or Post to make a request, and then use the `Response` object to get the response body, headers, status code, etc.

### Example

    package main
    
    import (
        "fmt"
        "github.com/tonnytg/webreq"
    )
    
    func main() {
        response, err := webreq.Get("https://api.example.com/data")
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println("Response:", response)
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

    headers := webreq.NewHeaders(map[string]string{
        "Content-Type": "application/json",
        "Authorization": "Bearer your_token",
    })
    
    request := webreq.NewRequest("POST")
    request.SetURL("https://api.example.com/data")
    request.SetData([]byte(`{"key":"value"}`))
    request.SetHeaders(headers.Headers)
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
        "fmt"
        "github.com/tonnytg/webreq"
    )
    
    func main() {
        response, err := webreq.Get("https://api.example.com/data")
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println("Response:", response)
    }


### Post

Create a body data to send in request

    package main
    
    import (
        "fmt"
        "github.com/tonnytg/webreq"
    )
    
    func main() {
        data := map[string]string{"key": "value"}
        response, err := webreq.Post("https://api.example.com/data", data)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println("Response:", response)
    }
        


## Advange Usage


### Custom Headers

    headers := map[string]string{
        "Content-Type": "application/json",
        "Authorization": "Bearer your_token",
    }
    
    response, err := webreq.GetWithHeaders("https://api.example.com/data", headers)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Response:", response)



### Erro Handlin

    response, err := webreq.Get("https://api.example.com/data")
    if err != nil {
        switch err.(type) {
        case webreq.HTTPError:
            fmt.Println("HTTP error occurred:", err)
        default:
            fmt.Println("An error occurred:", err)
        }
        return
    }
    fmt.Println("Response:", response)


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
