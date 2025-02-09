package main

import (
	"fmt";
)

func sumaDosValores(num1 int, num2 int) int {
	return num1 + num2;
}

func restaDosValores(num1 int, num2 int) int {
	return num1 - num2;
}

func mostrarSuma(sumaDosValores func(int, int) int, num1 int, num2 int) {
	fmt.Println("El resultado de la suma de ", num1, " y ", num2, " es: ", sumaDosValores(num1, num2));
}

func mostrarResta(sumaDosValores func(int, int) int, num1 int, num2 int) {
	fmt.Println("El resultado de la resta de ", num1, " y ", num2, " es: ", restaDosValores(num1, num2));
}

func main() {
	/*
	FUNCIONES EN GOLANG
	*/
	fmt.Println("FUNCIONES EN GOLANG");
	fmt.Println("Se capturan dos valores y se suma");

	var num1 int;
	var num2 int;

	fmt.Println("Ingrese el primer valor: ");
	fmt.Scan(&num1);

	fmt.Println("Ingrese el segundo valor: ");
	fmt.Scan(&num2);

	mostrarSuma(sumaDosValores, num1, num2);
	mostrarResta(restaDosValores, num1, num2);
}
