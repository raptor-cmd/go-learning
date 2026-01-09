package main

import (
	"time"
)

func main() {
	canal1 := make(chan string)
	canal2 := make(chan string)
	canal3 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		canal1 <- "Mensaje uno"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		canal2 <- "Mensaje dos"
	}()

	go func() {
		mensajeCanal := <-canal1
		time.Sleep(3 * time.Second)
		canal3 <- "Recibido en canal 3 mensaje de canal 1: " + mensajeCanal
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-canal1:
			println("Recibido del canal 1:", msg1)
		case msg2 := <-canal2:
			println("Recibido del canal 2:", msg2)
		case msg3 := <-canal3:
			println("Recibido del canal 3:", msg3)
		}
	}
}
