package main

import (
  "fmt"
  _ "hello/hello"
  _ "strconv"
  _ "strings"
)

func main() {
  m := map[string]int{
    "a": 100,
    "b": 200,
    "c": 300,
  }

  for k, v := range m {
    fmt.Println(k + ":", v)
  }
}
