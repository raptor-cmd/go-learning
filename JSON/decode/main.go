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
	jsonString := `{"name":"Juan","age":30,"email":"juan@example.com"}`

	var person Persona
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Nombre: %s, Edad: %d, Email: %s\n", person.Name, person.Age, person.Email)
}
