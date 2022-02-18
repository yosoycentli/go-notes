package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

//  El asterisco antes del tipo pc indica que se va a acceder a la dirección
//   en memoria donde están los valores de pc y ahí se modificará
func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2

}

func main() {

	a := 50
	// "&" significa que va a solicitar la dirección en memoria
	//  de la variable que le siga, en este caso la variable "a"
	b := &a

	fmt.Println(b)
	// "*" el asterico significa que va a apuntar o accederá hacia donde
	//  está guardado en memoria la variable que le siga, en este caso "b"
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)

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
