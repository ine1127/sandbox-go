package main

import (
  "fmt"
  "hello/hello"
  "strconv"
)

func main() {
  x := hello.Input("type a price")
  n, err := strconv.Atoi(x)

  if err != nil {
    fmt.Println("ERROR!")
    return
  }

  fmt.Print(x + "は、")

  if n % 2 == 0 {
    fmt.Println("偶数です。")
  } else {
    fmt.Println("奇数です。")
  }
}
