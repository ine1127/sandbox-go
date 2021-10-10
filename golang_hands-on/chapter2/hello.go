package main

import (
  "fmt"
  _ "hello/hello"
  _ "strconv"
  _ "strings"
)

func main() {
  m := []string{
    "one", "two", "three",
  }

  fmt.Println(m)
  m = insert(m, "*", 2)
  m = insert(m, "*", 1)
  fmt.Println(m)
}

func insert(a []string, v string, p int) (s []string) {
  s = append(a, "")
  s = append(s[:p+1], s[p:len(s)-1]...)
  s[p] = v
  return
}
