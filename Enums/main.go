package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateNames = map[ServerState]string{
	StateIdle:      "Inactivo",
	StateConnected: "Conectado",
	StateError:     "Error",
	StateRetrying:  "Retrying",
}

func (estadoServer ServerState) String() string {
	return stateNames[estadoServer]
}

func main() {

	redServidor := verificacionDeRed(StateIdle)

	fmt.Println("Estado del servidor:", redServidor)

	segundaRevisión := verificacionDeRed(redServidor)

	fmt.Println("Estado del servidor después de la segunda verificación:", segundaRevisión)

}

func verificacionDeRed(servidor ServerState) ServerState {
	switch servidor {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("Estado desconocido: %s", servidor))
	}
}
