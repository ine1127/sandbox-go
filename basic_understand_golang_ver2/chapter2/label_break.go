package main

import "fmt"

func main() {
LBL:
  for i := 0; i = 5; i++ {
    switch {
    case i == 3:
      break LBL
    default:
      fmt.Println(i)
    }
  }
}
