package main

import "fmt"

func main() {
  c := make(chan int, 10)

  c <- 0

  fmt.Println("cap:", cap(c))
  fmt.Println("len:", len(c))
}
