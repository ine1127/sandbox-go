package main

import (
  "fmt"
  "os"
)

func main() {
  os.MkdirAll("testdir", 0777)

  err := os.Rename("testdir", "newdir")

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  file, _ := os.Create("testfile")
  file.Close()

  err = os.Rename("testfile", "newfile")

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
