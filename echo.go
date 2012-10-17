package main

import (
    "fmt"
    "net/http"
    "os"
    "go-echoserver/src/metric"
)

func handler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path[1:] != "favicon.ico" {
      metric.Track("view")
    }
    fmt.Printf("[%s] ", r.RemoteAddr)
    fmt.Printf("%s ", r.Method)
    fmt.Printf("%s\n", r.URL.Path)

    fmt.Fprintf(w, "%s\n", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    port := os.Getenv("PORT")
    fmt.Println("listening on port " + port)
    http.ListenAndServe(":"+port, nil)
}
