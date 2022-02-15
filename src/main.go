package main

import "fmt"

// Functions

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	// Contants declaration
	const pi float64 = 3.1416
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicacion
	result = y * x
	fmt.Println("Multiplicación:", result)

	// Division
	result = y / x
	fmt.Println("Division", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental", x)

	// Decremental
	x--
	fmt.Println("Decremental", x)

	// declaración de variables

	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)
	// Le agrega el salto de línea a cada mensaje

	// Prinf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	// %s es por string y %d es para poner un entero, sino sabes
	// que clase de dato tendrás, debes usar %v
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	// Genera un string pero no lo imprime en consola, solo lo guarda como un string
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// tipo de dato
	// %T te dice el tipo de variable
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

	normalFunction("Hola Mundo")

	tripleArgument(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Value:", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("value1:", value1)
	fmt.Println("value2:", value2)

	// If we don't want one return value we must use the floor: "_" for
	//  the value we don't want to use

	value3, _ := doubleReturn(5)
	fmt.Println("value3:", value3)

}

// To ejecute this file you need to compile it in the terminal, we have two options:
//  1. the fast: go run src/main.go
//  2. the "hard" but with better performance:
// 	2.1 go build src/main.go
// 	2.2 ./main
