/*
- Exercício:
    - Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados
*/
package main 

import "fmt"

func retornaInt() (int) {
	return 10
}

func retornaIntString() ( int,  string) {
	return 1, "\nPrimero retorna um int e logo depois uma string"
}

func main() {
	fmt.Println(retornaInt())
	fmt.Println(retornaIntString())
}