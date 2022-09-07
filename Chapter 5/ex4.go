/*
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/
package main

import "fmt"

type food struct {
	nome string
	sabor string
	ondetem []string
	vaibemcom map[string]string
}

func main() {
	comida := food {
		nome: "Stroopwafel",
		sabor: "doce",
		ondetem: []string{"Na holanda", "na casa do seu tio rico"},
		vaibemcom: map[string]string{
			"no café da manhã": "chá",
			"no almoço": "cafezinho",
		},
	}

	fmt.Println(comida)
}