package main

import (
  "log"
  "os"
)

func main() {
  _, err := os.OpenFile("test", os.O_RDONLY, 0)

  if err != nil {
    log.Fatalln("エラー", err)
  }
}
