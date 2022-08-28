/*
- Partindo do código abaixo, utilize marshal para transformar []user em JSON.
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	primeiro, err := json.Marshal(u1)

	if err != nil {
		log.Fatal(err)
	}

	segundo, err := json.Marshal(u2)

	if err != nil {
		log.Fatal(err)
	}

	terceiro, err := json.Marshal(u2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(primeiro))
	fmt.Println(bytes.NewBuffer(segundo))
	fmt.Println(bytes.NewBuffer(terceiro))
}
