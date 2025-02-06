package main

import "fmt"

func main() {
	fmt.Println("PRograma para estudiar condicionales en Golang")

	var edad int = 18

	if edad >= 18 {
		fmt.Println("Esta persona es mayor de edad")
	} else {
		fmt.Println("Esta persona es menor de edad")
	}

	fmt.Println("------------------------------------------------------------------------------------------")

	var edadPersona int = 89

	if edadPersona >= 65 {
		fmt.Println("Esta persona es un adulto mayor")
	} else if edadPersona >= 40 {
		fmt.Println("Esta persona es un adulto")
	} else if edadPersona >= 18 {
		fmt.Println("Esta persona es un joven")
	} else if edadPersona >= 0 {
		fmt.Println("Esta persona es un niño")
	} else {
		fmt.Println("Esta persona no existe, edad incorrecta")
	}

	fmt.Println("Has terminado de estudiar condicionales en Golang")
}
