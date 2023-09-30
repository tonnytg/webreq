# Web Request for Go Packages
[![Go Report Card](https://goreportcard.com/badge/github.com/tonnytg/webreq)](https://goreportcard.com/report/github.com/tonnytg/webreq) [![codecov](https://codecov.io/gh/tonnytg/webreq/branch/main/graph/badge.svg?token=PYI6QQKGTV)](https://codecov.io/gh/tonnytg/webreq) ![example workflow](https://github.com/tonnytg/webreq/actions/workflows/go.yml/badge.svg) 

## What is this?

This is module to help you make web requests in Go, it is a wrapper around the standard library's `http` package.
You can use Get or Post to make a request, and then use the `Response` object to get the response body, headers, status code, etc.

## Why

So many times I needed to make a request in an API and after convert body to struct, and every time I needed to configure http.Client and put headers and body. Thinking about "don't repeat your self" I believe this unified code can gain time build an easy way to make this request.

What do you think? Liked, set star in this project to help me to help others!


### Install

> go get github.com/tonnytg/webreq

### Get

Create a slice of headers to use and input in request 

	headers := webreq.NewHeaders()
	headers.Add("Content-Type", "application/json")

	request := webreq.Builder("GET")
	request.SetURL("https://610aa52552d56400176afebe.mockapi.io/api/v1/friendlist")

	_, err := request.Execute()
	if err != nil {
		t.Error(err)
	}


### Post

Create a body data to send in request

    ...
	f := Friend{
		CreatedAt:  time.Now(),
		Name:       "Tonny",
	}

	// convert f to bytes
	fBytes, err := json.Marshal(f)
	if err != nil {
		t.Error(err)
	}

	request := webreq.Builder("POST")
	request.SetURL("https://610aa52552d56400176afebe.mockapi.io/api/v1/friendlist")
	request.SetBody(fBytes)

	body, err := request.Execute()
	if err != nil {
		t.Error(err)
	}
    ...
    
