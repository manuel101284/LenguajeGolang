package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		En golang solo existe el ciclo for, pero puede ser utilizado de diferentes maneras.
	*/
	var suma int = 0

	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			suma = suma + i
		}
		// fmt.Println(i);
	}
	fmt.Println("La suma de los números impares del 0 al 100 es: ", suma)

	// For Each
	mapaPaises := map[string]string{
		"Colombia":  "Bogotá",
		"Perú":      "Lima",
		"Chile":     "Santiago",
		"Argentina": "Buenos Aires",
		"Venezuela": "Caracas",
		"Brasil":    "Brasilia",
		"Uruguay":   "Montevideo",
		"Paraguay":  "Asunción",
		"Bolivia":   "La Paz",
		"Ecuador":   "Quito",
	}

	for k, v := range mapaPaises {
		fmt.Println("La capital de", k, "es", v)
	}

	for llave, valor := range mapaPaises {
		fmt.Println("Ciudad capital en el país ", llave, "es", valor)
	}

	// Repita Mientras
	var fruit string = ""

	for fruit != "PAPAYA" {
		fmt.Println("Ingresa el nombre de una fruta")
		fmt.Scan(&fruit)
		fruit = strings.ToUpper(fruit)
	}
}
