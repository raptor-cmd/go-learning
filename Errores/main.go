package main

import (
	"errors"
	"fmt"
)

var ErrorDeCafe = fmt.Errorf("No hay café")
var ErrorDeEnergia = errors.New("Ya no hay electricidad")

func hacerCafe(args int) error {
	if args == 2 {
		return ErrorDeCafe
	} else if args == 4 {
		return fmt.Errorf("Haciendo café: %w", ErrorDeEnergia)
	}
	return nil
}

func main() {
	for i := range 5 {
		if err := hacerCafe(i); err != nil {
			if errors.Is(err, ErrorDeCafe) {
				fmt.Println(err)
				fmt.Println("Por favor, trae más café")
			} else if errors.Is(err, ErrorDeEnergia) {
				fmt.Println(err)
				fmt.Println("Ahora no puedo calentar el agua")
			} else {
				fmt.Println("Error desconocido:", err)
			}
			continue
		}
		fmt.Println("Ya hay café")
	}
}
