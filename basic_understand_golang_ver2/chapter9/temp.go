package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

func main() {
  dir := os.TempDir()

  file, _ := ioutil.TempFile(dir, "test")

  fmt.Println(file.Name())
}
