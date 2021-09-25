package main

import (
  "fmt"
  "net"
  "os"
)

func main() {
  addrs, err := net.LookupHost("www.google.com")

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Printf("%q\n", addrs)
}
