package main

import (
  "fmt"
  "hello/hello"
  "strconv"
)

func main() {
  x := hello.Input("type a price")
  fmt.Print(x + "は、")

  if n, err := strconv.Atoi(x); err == nil {
    if n % 2 == 0 {
      fmt.Println("偶数です。")
    } else {
      fmt.Println("奇数です。")
    }
  } else {
    fmt.Println("整数ではありません。")
  }
}
