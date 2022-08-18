/*
- Crie um map com key tipo string e value tipo []string.
- Key deve conter nomes no formato sobrenome_nome
- Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
*/
package main

import ("fmt")

func main() {
	x := map[string][]string {
		"José_sumido": {"comer ratinhos", "ir pra rua"},
		"Argus_guardinha": {"pedir carinho", "ficar quietinho"},
		"Gudi_sonsa": {"Pedir comida", "miar por comida"},
	}

	for k, v := range x {
		fmt.Println(k[0])
		for i, item := range v {
			fmt.Println("\t", i, "-", item)
		}
	}
}