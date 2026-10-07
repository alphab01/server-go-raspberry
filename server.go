package main

import (
  "fmt"
  "net/http"
)

func h(w http.ResponseWriter, r *http.Request) {
  if (r.Method == "GET") {
    fmt.Println("get")
    w.Write([]byte("test"))
  }
}

func main() {
  http.HandleFunc("/", h)
  err := http.ListenAndServe(":8080", nil)
  if (err != nil) {
    fmt.Println(err)
  }
}
