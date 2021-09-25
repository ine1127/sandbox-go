package main

import (
  "fmt"
  "math"
)

func main() {
  for i := 1; i <= 5; i++ {
    fmt.Println(i, math.Sqrt(float64(i)))
  }
}
