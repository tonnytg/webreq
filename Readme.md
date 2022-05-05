# Package Web Request Request
# pkggowebreq

> go get -u github.com/tonnytg/pkggowebreq

Create a slice of headers to use and input in request 


    func main() {
        println("Hello, world.")
        url := "https://610aa52552d56400176afebe.mockapi.io/api/v1/friendlist"
        timeOut := 20
        headers := []string{}
        body, err := web.Get(url, headers, timeOut)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%s", body)
    }