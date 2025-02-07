package main

import "fmt"

func main() {
	var fruta string

	fmt.Println("Ingresa el nombre de una fruta")

	fmt.Scan(&fruta)

	switch fruta {
	case "Manzana":
		fmt.Println("Es una manzana")
	case "Pera":
		fmt.Println("Es una pera")
	case "Uva":
		fmt.Println("Es una uva")
	case "Mango":
		fmt.Println("Es un mango")
	default:
		fmt.Println("No está en la lista o no es una fruta")
	}
}
