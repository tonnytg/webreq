# Web Request for Go Packages

## What is this?

This is module to help you make web requests in Go, it is a wrapper around the standard library's `http` package.
You can use Get or Post to make a request, and then use the `Response` object to get the response body, headers, status code, etc.

### Install

> go get github.com/tonnytg/webreq

### Get

Create a slice of headers to use and input in request 

    println("Hello, world.")
    url := "https://www.google.com/robots.txt"
    timeOut := 20
    headers := map[string]string{
        "Content-Type": "application/json",
    }
    body, err := webreq.Get(url, headers, timeOut)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s\n", body)


### Post

Create a body data to send in request

    println("Hello, world.")
    url := "https://www.google.com/robots.txt"
    timeOut := 20
    headers := map[string]string{
        "Content-Type": "application/json",
    }
    bodyData := map[string]string{
        "name": "tonnytg",
        "age":  "20",
    }
    body, err := webreq.Post(url, headers, timeOut, bodyData)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s\n", body)
