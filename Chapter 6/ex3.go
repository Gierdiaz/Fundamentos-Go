/*
- Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
*/
package main

import ("fmt")

func main() {
	defer fmt.Println("Executa isso aqui depois da função")

	function()
}

func function() {
	fmt.Println("executa isso aqui primeiro e depois o defer")
}