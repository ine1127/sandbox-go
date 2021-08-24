package main

import "fmt"
import "os"

func main() {
  file, err := os.Open("test.txt")

  if err != nil {
    fmt.Println(err.Error())

    os.Exit(1)
  }

  file.Close()
  fmt.Println("OK")
}
