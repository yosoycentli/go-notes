package main

import (
	pk "curso_golang_platzi/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	// No se puede acceder a esta clase porque es de acceso privado
	// var myOtherCar pk.carPrivate
	// fmt.Println(carPrivate)

	pk.PrintMessage("Hola Platzi")
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
