package main

import (
  "fmt"
  _ "hello/hello"
  _ "strconv"
  _ "strings"
)

func main() {
  a := [5]int{10, 20, 30, 40, 50}
  b := a[0:3]
  fmt.Println(a)
  fmt.Println(b)
}
