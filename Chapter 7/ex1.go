/*
- Crie um valor e atribua-o a uma variável.
- Demonstre o endereço deste valor na memória.
*/
package main

import ("fmt")

func main() {
	x := 0

	y := &x

	fmt.Println("%T", x)
	fmt.Println("%T", y)

}

