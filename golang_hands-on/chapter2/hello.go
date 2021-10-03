package main

import (
  "fmt"
//  "hello/hello"
//  "strconv"
)

func main() {
  x := 5
  switch x {
  case f(1):
    fmt.Println("* first case. *")
  case f(2):
    fmt.Println("* sencond case. *")
  case f(3):
    fmt.Println("* third case. *")
  default:
    fmt.Println("* default case. *")
  }
}

func f(n int) int {
  fmt.Println("No,", n)
  return n
}
