package main

import (
  "encoding/base64"
  "fmt"
)

func main() {
  data := []byte{0x00, 0x01, 0x02, 0x03, 0x04}

  enc := base64.StdEncoding.EncodeToString(data)

  fmt.Println(enc)

  dec, _ := base64.StdEncoding.DecodeString(enc)

  fmt.Printf("% x\n", dec)
}
