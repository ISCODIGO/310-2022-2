package main

import (
	"fmt"
)

func main() {
	// Capturar datos de la consola
	var entero int
	var flotante float32
	var booleano bool
	var cadena string
	var p_cadena *string

	fmt.Print("Escriba un numero: ")
	//fmt.Scan(&entero)
	// Muestra los valores por defecto de varios tipos
	// En el caso de un puntero el valor por defecto es nil, lo que significa ese puntero no esta haciendo referencia a ninguna posicion de memoria
	fmt.Println("El valor es", entero, flotante, booleano, "[" + cadena + "]", p_cadena)

	// Declaracion previa
	cadena = "hola"
	// Declaracion y asignacion
	otra_cadena := "mundo"
	// El puntero accede la posicion de memoria de la variable
	p_cadena = &otra_cadena

	fmt.Println("Posicion de memoria de la variable", &otra_cadena)
	fmt.Println("Posicion de memoria del puntero", &p_cadena)
	fmt.Println("Que tiene almacenado el puntero", p_cadena)
	fmt.Println("Desreferencia del puntero", *p_cadena)
	*p_cadena = "adios"
	fmt.Println("Que hay ahora dentro de la variable", otra_cadena)

	var arreglo [5]int
	fmt.Println("Arreglo:", arreglo)

	// var p_arreglo *[5]int = &arreglo
	p_arreglo := &arreglo
	(*p_arreglo)[1] = 1
	fmt.Println("Arreglo modificado:", arreglo)
	fmt.Println("Que almacena el puntero de [5]int:", p_arreglo)

}

/*

Direccion		Variable		Valor asignado
10260			otra_cadena		"mundo"
0e030			p_cadena		10260

*/