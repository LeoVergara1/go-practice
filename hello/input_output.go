package main

import (
	"fmt"
)

func main() {

	edad := int(22)
	fmt.Println("Mi edad es: ", edad)
	bander := false
	fmt.Println("Mi bandera es: %s\n ", bander)
	nombre := "Juan"
	fmt.Printf("Mi nombre es: %s\n ", nombre)

	fmt.Scanf("%d", &edad)
	var nombreDos string
	fmt.Println("Coloca tú nombre")
	fmt.Scanf("%s", &nombreDos)
	fmt.Println("Mi nombre es: ", nombreDos)

}
