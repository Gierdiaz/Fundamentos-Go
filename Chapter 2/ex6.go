/*
- Crie um programa que:
- Atribua um valor int a uma variável
- Demonstre este valor em decimal, binário e hexadecimal
- Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
- Demonstre esta outra variável em decimal, binário e hexadecimal
*/
package main

import ("fmt")

const x int = 200

func main() {
  fmt.Printf("%d\t %b\t #%x\n", x, x, x)
  y := x << 1
  fmt.Printf("%d\t %b\t #%x", y, y, y)

}