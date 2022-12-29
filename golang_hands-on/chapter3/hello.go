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
	// この時点でchanに値は設定できない。
	// chanは複数スレッド間で値をやり取りするためのものであるが、
	// 送信側、受信側の双方で値の準備が整っていないといけないため。
	// このため、Goroutineによるスレッドを実行した後でないとchanは使えない
  // なので、このプログラムは失敗する
	c <- 100
	go total(c)
	time.Sleep(100 * time.Millisecond)
}
