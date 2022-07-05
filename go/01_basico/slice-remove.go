/*
El slice no tiene un metodo nativo para eliminar elementos. Es por eso que se 
define una funcion que realiza dicha operacion.
*/
package main

import "fmt"

func main() {
	var miSlice []int = []int{10, 20, 30, 40, 50}
	var eliminado int

	// [ini:fin] => a partir del 'ini' se toman fin - ini elementos
	fmt.Println(miSlice[4:5])

	fmt.Println(miSlice[:3])
	fmt.Println(miSlice[3:])

	miSlice, eliminado = remove(miSlice, 2)
	fmt.Println("Slice actualizado: ", miSlice)
	fmt.Println("Eliminado: ", eliminado)

	miSlice, eliminado = remove(miSlice, 2)
	fmt.Println("Slice actualizado: ", miSlice)
	fmt.Println("Eliminado: ", eliminado)

	var otro []int = []int{31, 331, 3433, 454, 2334}
	
	fmt.Println("Append: ", append(otro, 1000))
	fmt.Println("Como queda: ", otro)

}

func remove(s []int, posicion int) ([]int, int) {
	valor := s[posicion]
	return append(s[:posicion], s[posicion + 1:]...), valor
}