package main

import (
  "fmt"
  "time"
)

func main() {
  loc, _ := time.LoadLocation("Local")

  time := time.Date(2001, 2, 3, 11, 22, 33, 44, loc)

  fmt.Println(time)
}
