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

func h2(w http.ResponseWriter, r *http.Request) {
  if (r.Method == "GET") {
    fmt.Printn("get2")
    w.Write([]byte("hello"))
  }
}

func main() {
  http.HandleFunc("/", h)
  http.HandleFunc("/hi", h2)
  err := http.ListenAndServe(":8080", nil)
  if (err != nil) {
    fmt.Println(err)
  }
}
