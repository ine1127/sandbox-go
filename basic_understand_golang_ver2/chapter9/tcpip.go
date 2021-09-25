package main

import (
  "fmt"
  "io/ioutil"
  "net"
  "os"
)

func main() {
  conn, err := net.Dial("tcp", "golang.jp:80")

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  defer func() {
    conn.Close()
  }()

  fmt.Fprintf(conn, "GET /hello.html HTTP/1.1\r\n")
  fmt.Fprintf(conn, "HOST: golang.jp\r\n")
  fmt.Fprintf(conn, "\r\n")

  response, err := ioutil.ReadAll(conn)

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Println(string(response))
}
