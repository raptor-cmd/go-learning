package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numero int = 10
	numero2 := 20
	fmt.Println(numero, numero2)

	var numeroEntero = 10
	var numeroDoble = 20.5
	resultado := numeroEntero + int(numeroDoble)
	fmt.Println(resultado)

	var nombre string = "Josué"
	apellido := "Merino"
	nombreCompleto := nombre + " " + apellido
	fmt.Println(nombreCompleto)

	var edad int = 24
	identificacion := nombreCompleto + " " + fmt.Sprintf("%d", edad)
	identificacionOM := nombreCompleto + " " + strconv.Itoa(edad)
	fmt.Println(identificacion)
	fmt.Println(identificacionOM)
	fmt.Println(nombreCompleto, edad)
}
