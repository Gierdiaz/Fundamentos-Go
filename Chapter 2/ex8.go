/*
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.
*/
package main

import ("fmt")

const (
	firstYear = 1997 + iota
	secondYear
	thirdYear
	fourthYear
)

func main() {
	fmt.Println(firstYear, secondYear, thirdYear, fourthYear)
}