package mypackage

import "fmt"

// carPublic Car con acceso público (inician con mayúsculas)
type CarPublic struct {
	Brand string
	Year  int
}

// carPrivate de acceso privado (inician con minúsculas)
type carPrivate struct {
	brand string
	year  int
}

// PrintMessage imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}
