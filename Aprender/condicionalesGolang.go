package main

import "fmt"

func main() {
	fmt.Println("PRograma para estudiar condicionales en Golang")

	var edad int = 19

	fmt.Println("_----------------------------------UNA CONDICION----------------------------------__")

	if edad >= 18 {
		fmt.Println("Esta persona es mayor de edad")
	} else {
		fmt.Println("Esta persona es menor de edad")
	}

	fmt.Println("-----------------------UNA CONDICION MULTIPLES CASOS-----------------------------")

	var edadPersona int = 50

	var permiso bool = false

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

	fmt.Println("--------------------------------------DOS CONDICIONES---------------------------")

	if edadPersona >= 18 && permiso {
		fmt.Println("Esta persona es mayor de edad y tiene permiso")
	} else if edadPersona >= 18 && !permiso {
		fmt.Println("Esta persona es mayor de edad y no tiene permiso")
	} else if edadPersona < 18 {
		fmt.Println("Esta persona es menor de edad NO VIAJA SOLO")
	} else {
		fmt.Println("Esta persona no existe, edad incorrecta")
	}

	fmt.Println("Has terminado de estudiar condicionales en Golang")
}
