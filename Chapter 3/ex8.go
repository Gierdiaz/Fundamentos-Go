/*
- Crie um programa que utilize a declaração switch, sem switch statement especificado.
*/
package main

import ("fmt")

func main() {

	funcionario := "João Carlos"
	
	switch funcionario {
		case "João Carlos":
			fmt.Println("João Carlos está na team one")
		case "Camila Ferreiro":
			fmt.Println("Camila Ferriro está na team two")
		case "Patrick Fernandez":
			fmt.Println("Patrick Fernandez está na team three")
		default:
			fmt.Println("Essa pessoa não faz parte de nenhuma equipe")
	}
}