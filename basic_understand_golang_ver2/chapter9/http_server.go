package main

import (
  "fmt"
  "html"
  "net/http"
  "os"
)

func main() {
  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "<html><hody>Hello</body></html>")
  })

  http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
    val := r.FormValue("say")

    fmt.Fprintf(w, "<html><body>echo:%s</body></html>", html.EscapeString(val))
  })

  err := http.ListenAndServe(":8080", nil)

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
