package main

import "fmt"

func valoresMultiples(nombre1, nombre2, apellido string) (map[string]string, map[string]string) {
	mapa1 := make(map[string]string)
	mapa2 := make(map[string]string)

	mapa1[nombre1] = apellido
	mapa2[nombre2] = apellido

	return mapa1, mapa2
}

func nombreCompleto(nombre1, nombre2, apellido string) (string, string) {
	nombreCompleto1 := nombre1 + " " + apellido
	nombreCompleto2 := nombre2 + " " + apellido
	return nombreCompleto1, nombreCompleto2
}

func main() {

	nombre1 := "Josué"
	nombre2 := "Daniel"
	apellido := "Merino"

	mapaPrimerNombre, mapaSegundoNombre := valoresMultiples(nombre1, nombre2, apellido)

	fmt.Println("Mapa del primer nombre:", mapaPrimerNombre)
	fmt.Println("Mapa del segundo nombre:", mapaSegundoNombre)

	_, nombre := valoresMultiples(nombre1, nombre2, apellido)
	fmt.Println("Mapa del segundo nombre usando solo el segundo valor retornado:", nombre)

	completo1, completo2 := nombreCompleto(nombre1, nombre2, apellido)

	fmt.Println("Nombre completo 1:", completo1)
	fmt.Println("Nombre completo 2:", completo2)

}
