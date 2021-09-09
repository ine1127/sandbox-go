package main

import "os"

func main() {
  os.Mkdir("newdir", 0777)
  os.MkdirAll("a/b/c", 0777)
}
