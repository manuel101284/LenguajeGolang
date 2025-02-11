package main

import (
	"fmt"
)

var funciones = map[string]func(int, int) int{
	"SUMAR": func(num1 int, num2 int) int {
		return num1 + num2
	},
	"RESTAR": func(num1 int, num2 int) int {
		return num1 - num2
	},
	"MULTIPLICAR": func(num1 int, num2 int) int {
		return num1 * num2
	},
	"DIVIDIR": func(num1 int, num2 int) int {
		if num2 == 0 {
			fmt.Println("No es posible dividir por cero")
			return -999 // Se envia este valor puesto que el return debe ser un entero NO ES EL VALOR DE LA OPERACION
		}
		return num1 / num2
	},
}

func resultadoSuma(f func(int, int) int, num1 int, num2 int) {
	fmt.Println("El resultado de la suma de ", num1, " y ", num2, " es: ", f(num1, num2))
}

func resultadoResta(f func(int, int) int, num1 int, num2 int) {
	fmt.Println("El resultado de la resta de ", num1, " y ", num2, " es: ", f(num1, num2))
}

func presentarResultado(funcion string, num1 int, num2 int) {
	sumar := func(num1 int, num2 int) int {
		return num1 + num2
	}
	restar := func(num1 int, num2 int) int {
		return num1 - num2
	}
	resultado := 0

	if funcion == "sumar" {
		resultado = sumar(num1, num2)
	} else if funcion == "restar" {
		resultado = restar(num1, num2)
	}
	fmt.Println("El resultado al ", funcion, " ", num1, " y ", num2, " es: ", resultado)
}

func operarNumeros(operacion string, num1 int, num2 int) {
	f, exists:= funciones[operacion]

	if(!exists){
		fmt.Println("La operación ", operacion, " no existe")
		return
	}

	fmt.Println("El resultado de la ", operacion, " es: ", f(num1, num2))
}

func main() {
	var num1 int
	var num2 int

	fmt.Println("Ingrese el primer valor: ")
	fmt.Scan(&num1)
	fmt.Println("Ingrese el segundo valor: ")
	fmt.Scan(&num2)

	suma := func(num1 int, num2 int) int {
		return num1 + num2
	}
	resultadoSuma(suma, num1, num2)

	resta := func(num1 int, num2 int) int {
		return num1 - num2
	}
	resultadoResta(resta, num1, num2)
	presentarResultado("sumar", num1, num2)
	presentarResultado("restar", num1, num2)

	operarNumeros("SUMAR", num1, num2)
	operarNumeros("RESTAR", num1, num2)
	operarNumeros("MULTIPLICAR", num1, num2)
	operarNumeros("DIVIDIR", num1, num2)
	operarNumeros("POTENCIA", num1, num2)
}
