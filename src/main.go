package main

import "fmt"

func main() {
	// Array (Son inmutables)
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slice (Se puede modificar)
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// Array vacío
	slice2 := []int{}
	fmt.Println(slice2)
}

// To ejecute this file you need to compile it in the terminal, we have two options:
//  1. the fast: go run src/main.go
//  2. the "hard" but with better performance:
// 	2.1 go build src/main.go
// 	2.2 ./main
