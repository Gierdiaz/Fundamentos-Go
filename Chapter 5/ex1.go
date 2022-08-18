/*
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
- Nome
- Sobrenome
- Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
*/
package main

import ("fmt")

type pessoa struct {
	nome string
	sobrenome string
	saboresFavoritos string
}

func main() {
	cliente := pessoa {
		nome: "Fernanda",
		sobrenome: "Oliveira",
		saboresFavoritos: "Uva, Banana e Morango",
	}

	fmt.Println(cliente)
}