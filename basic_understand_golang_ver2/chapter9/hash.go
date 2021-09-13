package main

import (
  "crypto/md5"
  "crypto/sha1"
  "fmt"
  "io"
)

func main() {
  s := sha1.New()

  io.WriteString(s, "test")

  result1 := s.Sum(nil)

  fmt.Printf("SHA-1:% x\n", result1)

  m := md5.New()

  data := []byte{0x74, 0x65, 0x73, 0x74} //test

  m.Write(data)

  result2 := m.Sum(nil)

  fmt.Printf("MD5:% x\n", result2)
}
