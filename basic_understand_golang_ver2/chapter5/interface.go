package main

import "fmt"

type Calculator interface {
  Calculate(a int, b int) int
}

type Add struct {
}

func (x Add) Calculate(a int, b int) int {
  return a + b
}

type Sub struct {
}

func (x Sub) Calculate(a int, b int) int {
  return a - b
}

func main() {
  var add Add
  var sub Sub

  var cal Calculator

  cal = add

  fmt.Println("和:", cal.Calculate(1, 2))

  cal = sub

  fmt.Println("差:", cal.Calculate(1, 2))
}
