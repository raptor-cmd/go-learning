package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println("El valor de i es:", i)
		i++ // Comentario de una línea
		/* Comentario
		   de
		   varias
		   líneas */
	}

	fmt.Println("------------------------")
	for numero := 0; numero <= 3; numero++ {
		fmt.Println("El valor de numero es:", numero)
	}

	fmt.Println("------------------------")

	for rango := range 5 {
		fmt.Println("El valor de rango es:", rango)
	}

	fmt.Println("------------------------")

	for {
		fmt.Println("Bucle infinito")
		break
	}

	fmt.Println("------------------------")

	for valor := range 6 {
		if valor%2 == 0 {
			continue
		}

		fmt.Println("El valor es:", valor)
	}

	fmt.Println("------------------------")

	frutas := []string{"piña", "mango", "mandarina"}
	for fruta := range frutas {
		fmt.Printf("Valor: %s\n", frutas[fruta])
	}
}
