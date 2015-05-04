package main

import "fmt"

func main() {
  var n uint = 1
  var result = 23 << n // LEFT SHIFT 00010111 -> 00101110
  fmt.Println(result)
  result = 7 >> n // RIGHT SHIFT 0111 -> 0011
  fmt.Println(result)

  fmt.Println("Pointer: ", &n, "->", *&n)
  fmt.Println("Complex Númber: ", complex128(1))
}
