package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome      string  
	Idade     int64
	Profissao string
}

func main() {
	p1 := Pessoa {Nome: "Emmanuel", Idade: 25, Profissao: "Fotógrafo"}
	p2 := Pessoa {Nome: "João", Idade: 31, Profissao: "Cozinheiro"}

	e, err := json.Marshal(p1)

	if err != nil {
		fmt.Println(err)
	}

	j, err := json.Marshal(p2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(e))
	fmt.Println(string(j))
}