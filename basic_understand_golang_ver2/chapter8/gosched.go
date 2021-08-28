package main

import (
  "fmt"
  "os"
  "runtime"
)

func main() {
  go func() {
    fmt.Println("Goroutine")
    os.Exit(0)
  }()

  for {
    runtime.Gosched()
  }
}
