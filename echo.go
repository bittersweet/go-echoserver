package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    t := time.Now()
    fmt.Printf("%s ", t)
    fmt.Printf("[%s] ", r.RemoteAddr)
    fmt.Printf("%s ", r.Method)
    fmt.Printf("%s\n", r.URL.Path)

    fmt.Fprintf(w, "%s\n", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Printf("listening on port 8000\n")
    http.ListenAndServe(":8000", nil)
}
