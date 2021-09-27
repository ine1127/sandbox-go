package main

import "log"

func main() {
  log.Print("test", 1, 2, 3)
  log.Println("test", 1, 2, 3)
  log.Printf("test %d %d %d", 1, 2, 3)
}
