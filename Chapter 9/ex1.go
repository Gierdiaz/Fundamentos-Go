/*
- Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.

Parte de número 2
*/
package main

import (
	"fmt"
	"sync"
)

var waitgroup sync.WaitGroup

func main() {

	waitgroup.Add(10)


	go func () {
		
		for i := 0; i < 5; i++ {
			fmt.Println("Execução de número 1")
			waitgroup.Done()	
		}
	}()

	go func () {
		for i := 0; i < 4; i++ {
			fmt.Println("Execução de número 2")
			waitgroup.Done()
		}
		
	}()

	waitgroup.Wait()
}
