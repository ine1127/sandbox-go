package main

import (
  "strconv"
  "fmt"
)

func main() {
  f64, err := strconv.ParseFloat("123.456789", 32)
  fmt.Println(float32(f64), err)
}
