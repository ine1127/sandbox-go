package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  "os"
)

func main() {
  response, err := http.PostForm("http://golang.jp/hrllo.html",
    url.Values{
      "key1": {"Value1", "Value2"},
      "key2": {"Value"},
    },
  )

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
