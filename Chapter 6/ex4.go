/*
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
*/

package main

import ("fmt")

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func (p pessoa) Demonstrar() {
	fmt.Println("O nome completo dessa pessoa é:", p.nome,  p.sobrenome, "Essa pessoa tem ", p.idade, "anos")
}

func main() {
	gatinho := pessoa {
		nome: "José",
		sobrenome: "gatinho.",
		idade: 2,
	}

	gatinho.Demonstrar()
}