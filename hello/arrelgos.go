package main

import "fmt"

func main() {
	var arreglo [5]int

	fmt.Println("Arreglo inicial: ", arreglo)

	for i := 0; i < len(arreglo); i++ {
		arreglo[i] = i
	}
}
