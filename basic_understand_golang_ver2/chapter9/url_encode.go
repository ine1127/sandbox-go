package main

import (
  "fmt"
  "net/url"
)

func main() {
  data := "abcあいう/:"

  enc := url.QueryEscape(data)

  fmt.Println(enc)

  dec, _ := url.QueryUnescape(enc)

  fmt.Println(dec)
}
