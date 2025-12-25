package main

import (
	"fmt"
	"slices"
)

func main() {
	var arregloCadenas []string
	fmt.Println("datos:", "contenido", arregloCadenas, "condicion:", arregloCadenas == nil, "longitud:", len(arregloCadenas) == 0)

	arregloCadenas = make([]string, 3)
	arregloCadenas[0] = "a"
	arregloCadenas[1] = "b"
	arregloCadenas[2] = "c"
	fmt.Println("datos:", arregloCadenas)
	fmt.Println("datos específicos:", arregloCadenas[2])
	fmt.Println("longitud:", len(arregloCadenas))

	arregloCadenas = append(arregloCadenas, "d")
	arregloCadenas = append(arregloCadenas, "e", "f")
	fmt.Println("datos después de append:", arregloCadenas)
	fmt.Println("nueva longitud:", len(arregloCadenas))

	segundoArreglo := []string{"g", "h", "i"}
	fmt.Println("segundo arreglo:", segundoArreglo)

	tercerArreglo := []string{"g", "h", "i"}
	if slices.Equal(segundoArreglo, tercerArreglo) {
		fmt.Println("2 es igual que 3")
	}

	arregloCadenas = append(arregloCadenas, segundoArreglo...)
	arregloCadenas = append(arregloCadenas, tercerArreglo...)
	fmt.Println("Cadenas después de agregar segundo y tercer arreglo:", arregloCadenas)
}
