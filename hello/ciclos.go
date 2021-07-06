package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	a := 0
	for a < 10 {
		fmt.Println(a)
		if a == 3 {
			a++
			continue
		}
		if a == 5 {
			break
		}
		a++
	}
}
