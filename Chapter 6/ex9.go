/*
- Callback: passe uma função como argumento a outra função.
*/
package main

import ("fmt")

func main() {
	function2(function1)
}

func function1() {
	fmt.Println("Primeira função")
}

func function2(x func()) {
	fmt.Println("Segunda função (callback)")
	x()
}