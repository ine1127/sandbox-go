package main

import (
  "fmt"
  "os"
  "os/exec"
)

func main() {
  cmd := exec.Command("go", "help")

  result, err := cmd.Output()

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Printf("%s\n", result)
}
