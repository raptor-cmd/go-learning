package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	persona := Persona{
		Name:  "Juan",
		Age:   30,
		Email: "juan@example.com",
	}

	jsonData, err := json.Marshal(persona)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Tu archivo JSON es:", string(jsonData))
}
