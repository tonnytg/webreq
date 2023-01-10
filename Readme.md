# Web Request for Go Packages

## What is this?

This is module to help you make web requests in Go, it is a wrapper around the standard library's `http` package.
You can use Get or Post to make a request, and then use the `Response` object to get the response body, headers, status code, etc.

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
	fBytes, err := webreq.StructToBytes(f)
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
    
