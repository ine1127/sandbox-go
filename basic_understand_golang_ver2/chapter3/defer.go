package main

import "fmt"

func main() {
  fmt.Println("開始")

  defer delay()

  fmt.Println("終了")
}

func delay() {
  fmt.Println("delayが呼び出されました")
}
