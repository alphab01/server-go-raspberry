package main

import (
  "fmt"
  "io"
  "net/http"
)

func main() {
  r, err := http.Get("http://192.168.3.19:8080")
  if (err != nil) {
    fmt.Println(err)
  }
  defer r.Body.Close()
  d, err := io.ReadAll(r.Body)
  if (err != nil) {
    fmt.Println(err)
  }
  fmt.Printf("%s\n", d)
}
