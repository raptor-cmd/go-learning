package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("Escribe", i, "como")

	switch i {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	default:
		fmt.Println("un número desconocido")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("A descansar!!!")
	default:
		fmt.Println("Es un día laboral")
	}

	t := time.Now()
	fmt.Println("Son las", t)

	switch {
	case t.Hour() < 12:
		fmt.Println("Buenos días")
	case t.Hour() < 18:
		fmt.Println("Buenas tardes")
	default:
		fmt.Println("Buenas noches")
	}

	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Hoy es lunes")
	case time.Tuesday:
		fmt.Println("Hoy es martes")
	case time.Wednesday:
		fmt.Println("Hoy es miércoles")
	case time.Thursday:
		fmt.Println("Hoy es jueves")
	case time.Friday:
		fmt.Println("Hoy es viernes")
	case time.Saturday:
		fmt.Println("Hoy es sábado")
	case time.Sunday:
		fmt.Println("Hoy es domingo")
	}
}
