package main

import (
  "fmt"
  "time"
)

func hello(s string, n int) {
  for i := 1; i <= 10; i++ {
    fmt.Printf("<%d %s>", i, s)
    time.Sleep(time.Duration(n) * time.Millisecond)
  }
}

func main() {
  go hello("hello", 50)
  hello("bye!", 100)
}
