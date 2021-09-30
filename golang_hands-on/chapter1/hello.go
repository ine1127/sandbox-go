package main

import (
  "fmt"
  "hello/hello"
)

func main() {
  fmt.Println("Hello, Go-lang!")

  fmt.Print("123 * 45 = ")
  fmt.Println(123 * 45)

  name := hello.Input("type your name")
  fmt.Println("Hello, " + name + "!!")
}
