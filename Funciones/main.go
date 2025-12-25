package main

import "fmt"

func suma(numero1, numero2 int) int {
	resultado := numero1 + numero2
	return resultado
}

func sumaLarga(numero1, numero2, numero3 int) int {
	resultado := numero1 + numero2 + numero3
	return resultado
}

func resta(numero1, numero2 int) int {
	resultado := numero1 - numero2
	return resultado
}

func restaLarga(numero1, numero2, numero3 int) int {
	resultado := numero1 - numero2 - numero3
	return resultado
}

func multiplicacion(numero1, numero2 int) int {
	resultado := numero1 * numero2
	return resultado
}

func multiplicacionLarga(numero1, numero2, numero3 int) int {
	resultado := numero1 * numero2 * numero3
	return resultado
}

func division(numero1, numero2 int) int {

	if numero2 == 0 || numero1 == 0 {
		fmt.Println("Error: Division por cero no es permitida.")
		return 0
	}

	resultado := numero1 / numero2
	return resultado
}

func main() {

	var numero1, numero2, numero3 int

	fmt.Println("Ingrese el primer numero:")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese el segundo numero:")
	fmt.Scanln(&numero2)

	fmt.Println("Ingrese el tercer numero:")
	fmt.Scanln(&numero3)

	resultadoSuma := suma(numero1, numero2)
	fmt.Println("El resultado de la suma de los dos primeros numeros es:", resultadoSuma)

	resultadoSumaLarga := sumaLarga(numero1, numero2, numero3)
	fmt.Println("El resultado de la suma de los tres numeros es:", resultadoSumaLarga)

	resultadoResta := resta(numero1, numero2)
	fmt.Println("El resultado de la resta de los dos primeros numeros es:", resultadoResta)

	resultadoRestaLarga := restaLarga(numero1, numero2, numero3)
	fmt.Println("El resultado de la resta de los tres numeros es:", resultadoRestaLarga)

	resultadoMultiplicacion := multiplicacion(numero1, numero2)
	fmt.Println("El resultado de la multiplicacion de los dos primeros numeros es:", resultadoMultiplicacion)

	resultadoMultiplicacionLarga := multiplicacionLarga(numero1, numero2, numero3)
	fmt.Println("El resultado de la multiplicacion de los tres numeros es:", resultadoMultiplicacionLarga)

	resultadoDivision := division(numero1, numero2)
	fmt.Println("El resultado de la division de los dos primeros numeros es:", resultadoDivision)

}
