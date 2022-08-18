/*
- Usando uma literal composta:
     - Crie um array que suporte 5 valores to tipo int
     - Atribua valores aos seus índices
- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.
*/
package main

import ("fmt")

func main() {
	array := [5]string {"A", "B", "C", "D", "E"}

	for i, v := range array {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", array)

	
	
	
		
	
}