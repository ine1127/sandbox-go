package main

import "fmt"

func main() {
  for i := 0; i < 5; i++ {
    switch i {
    case 0:
      fmt.Println("0")
    case 1, 2:
      fmt.Println("1または2")
    default:
      fmt.Println("その他")
    }
  }
}
