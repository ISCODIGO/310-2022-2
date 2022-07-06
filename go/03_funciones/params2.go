/*
Ambas funciones trataran de aumentar en 2 la variable.
Sin embargo, solamente el paso por referencia mantiene 
*/
package main

import "fmt"

func main() {
	a := 10
	paso_valor(a)
	fmt.Println("Paso por valor despues: ", a)

	paso_referencia(&a)
	fmt.Println("Paso por referencia despues: ", a)
}

func paso_valor(a int) {
	a += 2
	fmt.Println("Paso por valor dentro de la funcion: ", a)
}

func paso_referencia(a *int) {
	*a += 2
	fmt.Println("Paso por referencia dentro de la funcion: ", *a)
}