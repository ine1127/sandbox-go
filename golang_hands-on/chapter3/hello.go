package main

import (
	"fmt"
	"strconv"
  "sync"
	"time"
)

// SrData is structure.
type SrData struct {
  msg string
  mux sync.Mutex
}

func main() {
  sd := SrData{msg: "Start"}
  prmsg := func(nm string, n int) {
    fmt.Println(nm, sd.msg)
    time.Sleep(time.Duration(n) * time.Millisecond)
  }

  main := func(n int) {
    const nm string = "*main"
    sd.mux.Lock() // lock

    for i := 0; i < 5; i++ {
      sd.msg += " m" + strconv.Itoa(i)
      prmsg(nm, 100)
    }
    sd.mux.Unlock() // unlock
  }

  hello := func(n int) {
    const nm string = "hello"
    sd.mux.Lock() // lock

    for i := 0; i < 5; i++ {
      sd.msg += " h" + strconv.Itoa(i)
      prmsg(nm, n)
    }
    sd.mux.Unlock() // unlock
  }

  go main(100)
  go hello(50)
  time.Sleep(5 * time.Second)
}
