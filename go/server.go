// This is a basic web server written in Go which serves it's own source code

package main

import (
    "encoding/base64"
    "fmt"
    "net/http"
)

var sourceCode string

func main() {
    fmt.Println("Server running on port 8080...")

    http.HandleFunc("/", webRoot)

    http.ListenAndServe(":8080", nil)
}

func webRoot(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, b64Decoder(sourceCode))
    fmt.Printf("%s \"%s %s %s\" \"%s\"\n", req.RemoteAddr, req.Method, req.URL, req.Proto, req.Header.Get("User-Agent"))
}

func b64Decoder(encodedString string) string {
    decodedByte, _ := base64.StdEncoding.DecodeString(encodedString)

    decodedString := string(decodedByte)

    return decodedString
}
