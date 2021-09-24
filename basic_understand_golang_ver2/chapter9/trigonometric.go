package main

import (
  "fmt"
  "math"
)

func main() {
  var degree float64 = 45

  radian := degree * math.Pi / 180

  fmt.Println("角度(ラジアン):", radian)

  fmt.Println("正弦:", math.Sin(radian))

  fmt.Println("余波:", math.Cos(radian))

  fmt.Println("正接:", math.Tan(radian))
}
