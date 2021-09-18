package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println(time.Now())

  time.Sleep(10 * time.Second)

  fmt.Println(time.Now())

  time.Sleep(1 * time.Minute)

  fmt.Println(time.Now())
}
