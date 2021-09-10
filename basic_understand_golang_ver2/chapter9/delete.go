package main

import "os"

func main() {
  os.MkdirAll("a/b/c/d", 0777)
  os.MkdirAll("x/y/z", 0777)

  os.Remove("a/b/c/d")
  os.RemoveAll("x")

  file, _:= os.Create("test")
  file.Close()

  os.Remove("test")
}
