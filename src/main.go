package main

import "fmt"

func main() {

	// For condicinal
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	// For reverse
	reverseCount := 10
	for reverseCount >= 0 {
		fmt.Println(reverseCount)
		reverseCount--
	}

}

// To ejecute this file you need to compile it in the terminal, we have two options:
//  1. the fast: go run src/main.go
//  2. the "hard" but with better performance:
// 	2.1 go build src/main.go
// 	2.2 ./main
