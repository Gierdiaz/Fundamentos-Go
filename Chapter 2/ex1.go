/*
Base - 10: decimal, 0-9
Base - 2: binário, 0-1
Base - 16: hexadecimal, 0-f
*/
package main

import ("fmt")

func main() {
  a := 100

  fmt.Printf("%d\t %b\t %#x", a, a, a)
}