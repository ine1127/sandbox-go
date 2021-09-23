package main

import (
  "fmt"
  "strconv"
)

func main() {
  var f64 float64 = 1.23456789

  fmt.Println(strconv.FormatFloat(f64, 'e', 8, 64))
  fmt.Println(strconv.FormatFloat(f64, 'f', 4, 64))

  var f32 float32 = 1.23

  fmt.Println(strconv.FormatFloat(float64(f32), 'e', 2, 32))
}
