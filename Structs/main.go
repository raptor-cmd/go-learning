package main

import "fmt"

type persona struct {
	nombre string
	edad   int
}

func nuevaPersona(nombre string, edad int) *persona {
	nuevoIndividuo := persona{nombre: nombre}
	nuevoIndividuo.edad = 42

	return &nuevoIndividuo
}

func main() {
	fmt.Println(persona{"Juan", 42})

	fmt.Println(persona{nombre: "Santi", edad: 34})

	fmt.Println(persona{nombre: "Ana"})

	fmt.Println(nuevaPersona("Felipe", 50))

	personita := persona{nombre: "Laura", edad: 28}
	fmt.Println("Nombre:", personita.nombre)

	edadPersonita := &personita
	fmt.Println(edadPersonita.edad)

	edadPersonita.edad = 5
	fmt.Println("Edad modificada:", edadPersonita.edad)
	fmt.Println(personita)
}
