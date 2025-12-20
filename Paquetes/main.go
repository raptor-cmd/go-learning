package main

import (
	"fmt"
	"os"
)

func main() {

	envVar := os.Getenv("HOME")
	if envVar == "" {
		fmt.Println("La variable de entorno HOME no está definida.")
	} else {
		fmt.Println("La variable de entorno HOME está definida como:", envVar)
	}

	file, err := os.Create("ejemplo.txt")

	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
	}

	defer file.Close()
}
