package main

import (
	"fmt"
	"time"
)

func worker(id int, tareas <-chan int, resultados chan int) {
	for i := range tareas {
		fmt.Printf("Worker %d: procesando tarea %d\n", id, i)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "Tarea iniciada", i)
		resultados <- i * 2
	}
}

func main() {
	const numeroTareas = 5

	tareas := make(chan int, numeroTareas)
	resultados := make(chan int, numeroTareas)

	for w := 1; w <= 3; w++ {
		go worker(w, tareas, resultados)
	}

	for i := 0; i < numeroTareas; i++ {
		tareas <- i
	}
	close(tareas)

	for a := 0; a < numeroTareas; a++ {
		mostrarResultado := <-resultados
		fmt.Println("Resultado:", mostrarResultado)
	}
}
