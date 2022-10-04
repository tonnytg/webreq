# Package Web Request Request
# pkggowebreq

> go get -u github.com/tonnytg/webreq

Create a slice of headers to use and input in request 


    func main() {
        println("Hello, world.")
        url := "https://www.google.com/robots.txt"
        timeOut := 20
        headers := []string{}
        body, err := web.Get(url, headers, timeOut)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%s", body)
    }
