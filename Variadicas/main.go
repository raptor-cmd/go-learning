package main

import "fmt"

func suma(numeros ...int) {
	fmt.Println(numeros, " ")
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Println("Total:", total)
}

func main() {

	suma(1, 2)
	suma(1, 2, 3)

	numeros := []int{1, 2, 3, 4, 5}
	suma(numeros...)

}
