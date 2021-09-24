package main

import (
  "fmt"
  "math"
)

func main() {
  val := math.Pi

  fmt.Println("元の値:", val)

  fmt.Println("切り上げ:", math.Ceil(val))

  fmt.Println("切り捨て:", math.Floor(val))

  val *= -1

  fmt.Println("元の値:", val)

  fmt.Println("切り上げ:", math.Ceil(val))

  fmt.Println("切り捨て:", math.Floor(val))
}
