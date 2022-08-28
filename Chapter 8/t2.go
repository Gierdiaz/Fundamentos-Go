package main

import ("fmt"
		"app/golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "1dsfdsf"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))
	
}