package main

import (
	"fmt"
	"maps"
)

func main() {
	mapa := make(map[string]int)
	mapa["Josue"] = 5
	mapa["Merino"] = 6

	fmt.Println("mapa:", mapa)

	version1 := mapa["Merino"]
	fmt.Println("version1:", version1)

	version2 := mapa["Josue"]
	fmt.Println("version2:", version2)

	fmt.Println("len(mapa):", len(mapa))

	_, dato := mapa["Josue"]
	fmt.Println("dato:", dato)

	delete(mapa, "Josue")
	fmt.Println("mapa despues de delete:", mapa)

	clear(mapa)
	fmt.Println("mapa despues de clear:", mapa)

	nuevoMapa1 := map[string]int{
		"Josue":  1,
		"Merino": 2,
	}
	fmt.Println("nuevoMapa1:", nuevoMapa1)

	nuevoMapa2 := map[string]int{
		"Josue":  1,
		"Merino": 2,
	}
	fmt.Println("nuevoMapa2:", nuevoMapa2)

	if maps.Equal(nuevoMapa1, nuevoMapa2) {
		fmt.Println("nuevoMapa1 y nuevoMapa2 son iguales")
	}
}
