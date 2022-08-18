/*
- Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
    - Crie um tipo para um struct chamado "pessoa"
    - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
    - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
    - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
    - Demonstre no seu código:
        - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
        - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
*/
package main

import ("fmt")

//struct pessoa
type pessoa struct {
	nome string
	sobrenome string
}

//método com receive pessoa
func (p *pessoa) falar() {
	fmt.Println("Olá,", p.nome, p.sobrenome)
}

//interface implementada por tipos com o método 
type humanos interface {
	falar()
}

func main() {
	persone := pessoa {
		nome: "José",
		sobrenome: "gatinho",
	}

	persone.falar() // <- É um shortcut para (&persone).falar()
	(&persone).falar() // <- PE a maneira "mais correta.""

	dizerAlgumaCoisa(&persone) // <- Funciona
	//dizerAlgumaCoisa(persone) // < -Não funciona
}



func dizerAlgumaCoisa(h humanos) {
	h.falar() 
}