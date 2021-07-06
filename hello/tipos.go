package main

import (
	"strconv"
)

func main() {
	edad := 23

	edad_str := strconv.Itoa(edad)

	edad_int, _ := strconv.Atoi(edad_str)

	println(edad_str)
	println(edad_int)

}
