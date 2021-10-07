package main

import (
  "fmt"
  _ "hello/hello"
  _ "strconv"
  _ "strings"
)

func main() {
  a := [3]int{10, 20, 30}
  b := a[0:2]
  fmt.Println(a)
  fmt.Println(b)

  b = append(b, 1000)
  fmt.Println(a)
  fmt.Println(b)

  b = append(b, 1000)
  fmt.Println(a)
  fmt.Println(b)
}
