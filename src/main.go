package main

import "fmt"

type car struct {
	brand string
	year  int
	model string
	color string
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

	// Ejercicio
	var newCar car
	newCar.brand = "Volks Wagen"
	newCar.color = "Gray"
	newCar.model = "Vento"
	newCar.year = 2015

	fmt.Println("New Car:", newCar)

}

// To ejecute this file you need to compile it in the terminal, we have two options:
//  1. the fast: go run src/main.go
//  2. the "hard" but with better performance:
// 	2.1 go build src/main.go
// 	2.2 ./main
