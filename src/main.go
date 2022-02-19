package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	// Instanciar echo
	e := echo.New()

	// Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
	e.Logger.Fatal(e.Start(":1323"))
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
