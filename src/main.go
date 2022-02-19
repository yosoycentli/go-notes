package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	// el número establece la capacidad máxima de goroutines que soporta el channel " c"
	// en este caso es 2, pero si no le pones nada se creará la capacidad de forma dinámica
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	// len nos dice la cantidad de goroutines que hay en un channel (2)
	// cap nos dice la capacidad de goroutines que soporta el channel (2)
	fmt.Println(len(c), cap(c))

	// Range y close

	// close cierra el canal
	close(c)

	// este mensaje ya no puede entrar porque el canal se cerró, además de que ya está lleno
	//c <- "Mensaje3"

	// el for nos ayuda a hacer un recorrido por los mensajes del canal c
	for message := range c {
		fmt.Println(message)
	}

	// Select

	// cuando tenemos varios canales y no tenemos certeza de cual va a responder primero
	//  usamos select

	// Como no tienen especificado la capacidad, Go la asignará de forma dinñamica
	email1 := make(chan string)
	email2 := make(chan string)
	go message("mensaje1", email1)
	go message("mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
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
