package main

import (
    "fmt"
    "net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json; charset=utf-8")
    fmt.Fprintf(writer, "{\"message\": \"Hello, world!\"}")
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)
}
