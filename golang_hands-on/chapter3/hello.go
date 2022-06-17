package main

import (
  "fmt"
  // "time"
)

func total(n int, c chan int) {
  t := 0
  for i := 1; i <= n; i++ {
    t += 1
  }
  c <- t
}

func main() {
  c := make(chan int)
  go total(100, c)
  fmt.Println("total:", <-c)
}
