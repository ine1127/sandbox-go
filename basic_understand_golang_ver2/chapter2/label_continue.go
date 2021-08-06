package main

import "fmt"

func main() {
LBL:
  for i := 0; i < 3; i++ {
    fmt.Println(i)
    for j := 0; j < 3; j++ {
      fmt.Println("   ", j)
      if i == 1 && j == 1 {
        continue LBL
      }
    }
  }
}
