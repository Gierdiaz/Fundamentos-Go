/*
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
- "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/
package main

import ("fmt")

func main() {
	ss := [][]string{
		[]string{"José", "gatinho", "comer ratinho"},

		[]string{"Gudi","gatinha","comer lixo"},

		[]string{"Argus","gatinho","guardinha"},

	}

	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}
	}
}