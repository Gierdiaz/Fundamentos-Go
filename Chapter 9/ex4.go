/*
* Utilizando goroutines, crie um programa incrementador:
*   - Tenha uma variável com o valor da contagem
*   - Crie um monte de goroutines, onde cada uma deve:
*       - Ler o valor do contador
*       - Salvar este valor
*       - Fazer yield da thread com runtime.Gosched()
*       - Incrementar o valor salvo
*       - Copiar o novo valor para a variável inicial
*   - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
*   - Demonstre que há uma condição de corrida utilizando a flag -race
*
* Resolver a condição anterior utilizando Mutex
 */
 package main

 import (
	"fmt"
	"runtime"
	"sync"
)

var waitgroup sync.WaitGroup

var mu sync.Mutex

const quantidadeGoroutines = 100

var contador int

func main() {
	
	CriarRoutines(quantidadeGoroutines)

	waitgroup.Wait()

	fmt.Println("Total de goroutines:",quantidadeGoroutines)
	fmt.Println("Total de contador:",contador)

}

func CriarRoutines(i int) {
	waitgroup.Add(i)

	for j := 0; j < i; j++ {
		go func () {
			mu.Lock()
			variavel := contador
			runtime.Gosched()
			variavel ++
			contador = variavel
			mu.Unlock()
			waitgroup.Done()
		}()
		
	}

}
	
