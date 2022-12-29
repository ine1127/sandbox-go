package main

import (
	"fmt"
	"time"
)

func total(c chan int) {
	n := <-c
	fmt.Println("n = ", n)

	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total:", t)
}

func main() {
	c := make(chan int)
	// Goroutineを先に持ってくることでchanに値を設定することができる
	go total(c)
	c <- 100
	time.Sleep(100 * time.Millisecond)
}
