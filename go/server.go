// This is a basic web server written in Go which serves it's own source code

package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        http.ServeFile(res, req, "./server.go")
    })
    http.ListenAndServe(":8080", nil)
}
