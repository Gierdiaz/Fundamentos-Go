/*
- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
*/
package main

import "fmt"

func function(x... int) int {
	total := 0
	for _, v := range x {

		total = total + v
	}

	return total

}	

func function2( x[]int) int {
	total := 0
	for _,v := range x {
		total = total + v
	}

	return total
}

func main() {
	si := []int {1,2,3,4,5,6,7,8,9,10,20,30,40,50,60,70,80,90,100}
	
	fmt.Println(function(si...))
	fmt.Println(function2(si))
}