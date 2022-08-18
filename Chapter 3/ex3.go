/*
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/
package main

import ("fmt")

func main() {
	
	anoNascimento := 1997
	anoAtual := 2022

	for anoNascimento <= anoAtual {
		fmt.Println(anoNascimento)
		anoNascimento++
	}
}