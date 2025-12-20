package main

import "fmt"

func main() {

	nombre := "Josué"
	edad := 24

	if nombre == "Josué" {
		fmt.Println("Hola Josué")
	} else {
		fmt.Println("Hola desconocido")
	}

	if edad >= 18 {
		fmt.Println("Ya puedes votar")
	}

	if 9%2 == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}

	if numero := 99; numero < 0 {
		fmt.Println(numero, "es negativo")
	} else if numero < 10 {
		fmt.Println(numero, "es positivo")
	} else {
		fmt.Println(numero, "es un número mayor a 9")
	}
}
