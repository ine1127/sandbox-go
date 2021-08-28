package main

import (
  "fmt"
  "os"
)

func main() {
  go func() {
    fmt.Println("Goroutine")
    os.Exit(0)
  }()

  for {
  }
}
