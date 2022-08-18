/*
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
*/

package main

import ("fmt")

type pessoa struct {
	nome string
	sobrenome string
	saboresFavoritos []string
}

func main() {

	mapa := make(map[string]pessoa)

	mapa["Uva"] = pessoa {
		nome: "Luana",
		sobrenome: "Silva",
		saboresFavoritos: []string {"pistache", "baunilha", "abacaxi"}}
	



	for _, valor := range mapa {
		fmt.Println("Meu nome é", valor.nome, "e meus sabores favoritos são: ")
		for _, valor := range valor.saboresFavoritos {
			fmt.Println("-", valor)
		}
	}

	
}

