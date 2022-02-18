package main

import "fmt"

func main() {

	// Condición específica:
	//  Usar cuando vas a iterar sonbre una sola variable
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condicion:
	//  Utilizar cuando vas a hacer otra cosa.
	//  Como si una variable es mayor a otra o similares
	value := 50
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}
}

// To ejecute this file you need to compile it in the terminal, we have two options:
//  1. the fast: go run src/main.go
//  2. the "hard" but with better performance:
//  2.1 go build src/main.go
//  2.2 ./main
