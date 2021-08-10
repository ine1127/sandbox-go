package main

import "os"

func main() {
  file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0666)

  if err != nil {
    os.Exit(1)
  }

  defer file.Close()

  file.WriteString("あいうえお")
}
