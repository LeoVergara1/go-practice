package main

import "fmt"

func main() {
	arreglo := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Recuerda que el valor de una variable en go es cero por defecto
	slice := arreglo[0:5]                           // Partir un arreglo en una subarreglo, esto se llamara una slice
	println(arreglo)
	fmt.Println(arreglo)
	fmt.Println(arreglo[0:5])
	fmt.Println(slice)
}
