package main

import "fmt"

func main() {
  defer fmt.Println("defer")

  f1()
}

func f1() {
  panic("パニック発生")
}
