package main

import "fmt"

func main() {
  var ptr *int

  var i int = 12345

  ptr = &i

  fmt.Println("iのアドレス(&i):", &i)
  fmt.Println("ptrの値(変数iのアドレス(ptr)):", ptr)

  fmt.Println("iの値(i):", i)
  fmt.Println("ポインタ経由のiの値(*ptr):", *ptr)

  *ptr = 999

  fmt.Println("ポインタ経由で変更したiの値(i):", i)
}
