package main

import "fmt"

func main(){
	var cadena01 string = "Hola Mundo"
	var numeroint int = 18
	var edad int = 39

	nombre := "Ricardo";

	fmt.Println("Variable tipo string: ", cadena01);
	fmt.Println("Variable tipo int: ", numeroint);
	fmt.Println("Variable Edad: ", edad);
	fmt.Println("Mi nombre es: " + nombre);

	/*Variables numéricas*/
	var enteroPequeno8 int8 = -127;
	var enteroPequeno16 int16 = 32767;
	var numeroEntero int = 18;
	var numeroFlotante float32 = 3.14;
	var numeroFlotante64 float64 = 3.141592;
	var numeroComplejo complex64 = 3 + 5i;
	
	fmt.Println("Variable tipo int8: ", enteroPequeno8);
	fmt.Println("Variable tipo int16: ", enteroPequeno16);
	fmt.Println("Variable tipo int: ", numeroEntero);
	fmt.Println("Variable tipo float: ", numeroFlotante);
	fmt.Println("Variable tipo float64: ", numeroFlotante64);
	fmt.Println("Variable tipo complex: ", numeroComplejo);

	/*  ARREGLOS  */
	var listaFrutas = [5] string{"Banano", "Manzana", "Pera", "Uva", "Mango"};
	fmt.Println("Arreglo de frutas: ", listaFrutas);
	fmt.Println(listaFrutas[0]);
	fmt.Println(listaFrutas[2]);
	fmt.Println(listaFrutas[4]);

	listaPaises := [4] string{"Colombia", "Perú", "Argentina", "Ecuador"};
	fmt.Println("Arreglo de paises: ", listaPaises);
	listaPaises[0] = "Venezuela";
	fmt.Println("Arreglo de paises: ", listaPaises);

	/* SLICES */
	listaColores := [] string{"Rojo", "Azul", "Verde", "Amarillo", "Blanco"};
	fmt.Println("Lista de colores: ", listaColores);
	listaColores = append(listaColores, "Negro");
	fmt.Println("Lista de colores: ", listaColores);
	listaColores = append(listaColores, "Gris");
	fmt.Println("Lista de colores: ", listaColores);
	
	listaColores2 := listaColores[1:5];
	listaColores3 := listaColores[2:];
	listaColores4 := listaColores[:4];
	fmt.Println("Lista de colores: ", listaColores2);
	fmt.Println("Lista de colores: ", listaColores3);
	fmt.Println("Lista de colores: ", listaColores4);


	listaPrimos := [] int{2, 3, 5, 7, 11, 13, 17, 19, 23};
	fmt.Println("Lista de números primos: ", listaPrimos);
	
}
