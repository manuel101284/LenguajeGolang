/*
** Mapas - Diccionarios

	No necesitan una longitud determinada desde el inicio

**
*/
package main

import "fmt"

func main() {
	mapa01 := map[string]string{
		"Colombia":  "Bogota",
		"Peru":      "Lima",
		"Brasil":    "Brasilia",
		"Argentina": "Buenos Aires",
		"Venezuela": "Caracas",
	}

	fmt.Println("El mapa de paises es: ", mapa01)

	fmt.Println("La capital de Venezuela es: " + mapa01["Venezuela"])

	mapa01["Chile"] = "Santiago"

	fmt.Println("El mapa de paises es: ", mapa01)

	delete(mapa01, "Brasil")

	fmt.Println("El mapa de paises es: ", mapa01)

	fmt.Println("La capital de Brasil es: " + mapa01["Brasil"])
}
