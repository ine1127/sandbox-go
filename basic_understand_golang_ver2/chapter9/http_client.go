package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

func main() {
  response, err := http.Get("http://golang.jp/hello.html")

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Println("status:", response.Status)

  body, err := ioutil.ReadAll(response.Body)

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Println(string(body))
}
