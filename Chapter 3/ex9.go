/*
- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
*/
package main

import ("fmt")

func main() {
	esporteFavorito := "Wrestling"

	switch esporteFavorito {
		case "Wrestling":
			fmt.Println("Never stop wrestling")
		default:
			fmt.Print("Talvez o esporte digitado não esteja registrado")
	}
		

}