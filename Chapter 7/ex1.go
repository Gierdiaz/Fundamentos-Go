/*
- Crie um valor e atribua-o a uma variável.
- Demonstre o endereço deste valor na memória.
*/
package main

import ("fmt")

func main() {
	z := 0

	y := &z

	fmt.Println("%T", z)
	fmt.Println("%T", y)

}

