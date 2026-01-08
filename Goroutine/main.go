package main

import (
	"fmt"
	"time"
)

func function(from string) {

	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}

}

func main() {
	function("direct mode")

	go function("modo go routine")

	go func(mensaje string) {
		fmt.Println(mensaje)
	}("Enviando un mensaje")

	time.Sleep(1 * time.Second)
	fmt.Println("Fin del programa")

}
