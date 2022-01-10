package main

import (
  "net/http"
  "fmt"
)

func main() {
  http.HandleFunc("/", HelloServer)
  http.ListenAndServe(":80", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World!\nPath: %s", r.URL.Path[1:])
}
