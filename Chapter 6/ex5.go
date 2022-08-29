/*
- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura.
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
*/
package main

import (
	"fmt"
	"math"
)


type quadrado struct {
	lado float64
}

func (q quadrado) Area() {
	resultado := math.Pow(q.lado,2)
	fmt.Println(resultado)
}


type circulo struct {
	raio float64
}


func (c circulo) Area() {
	resultado := math.Pi*math.Pow(c.raio,2)
	fmt.Println(resultado)
}

type figura interface {
	Area()
}

func info(f figura) {
	f.Area()
}

func main() {
	x := quadrado {
		lado: 4,
	}

	y := circulo {
		raio: 10,
	}

	//x.Area()
	//y.Area()

	info(x)
	info(y)
}