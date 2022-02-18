package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	calcular(myCuadrado)
	calcular(myRectangulo)

	// Lista interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)

}

// To ejecute this file you need to compile it in the terminal, we have two options:
//  1. the fast: go run src/main.go
//  2. the "hard" but with better performance:
// 	2.1 go build src/main.go
// 	2.2 ./main

//  Si te sale el error:
//   go run src/main.go
//   src/main.go:4:2: package curso_golang_platzi/mypackage is not in GOROOT (/usr/local/Cellar/go/1.17.5/libexec/src/curso_golang_platzi/mypackage)

//  Debes correr el siguiente comando:
//   go mod init
//   Creará un archivo go.mod que te ayudará a correr el main.go de manera correcta
