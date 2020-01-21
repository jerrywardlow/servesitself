// This is a basic web server written in Go which serves it's own source code

package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Server running...")

    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        http.ServeFile(res, req, "./server.go")
        fmt.Printf("%s \"%s %s %s\" \"%s\"\n", req.RemoteAddr, req.Method, req.URL, req.Proto, req.Header.Get("User-Agent"))
    })

    http.ListenAndServe(":8080", nil)
}
