package greetings

import "fmt"

// Hola retorna un saludo para una persona especifica
func Hola(name string) string {
	message := fmt.Sprintf("Hola, %v. Bienvenido!", name)
	return message
}