package main

import "fmt"

func main() {
	// Un array -> tamaño fijo
	miArray := [...]int{20, 30, 5, 2, 1, 4, 10, 23, 45, 67}
	
	// Pudo declararse indicando el tamaño exacto:
	// miArray2 := [10]int{20, 30, 5, 2, 1, 4, 10, 23, 45, 67}
	
	// Tambien pudo declararse de la siguiente forma: [2, 0, 0, 0, 0]
	// miArray3 := [5]int{2}

	// Otra forma es utilizando make
	// miArray4 := make([]int, size, capacity)
	
	fmt.Printf("%v del tipo %T\n", miArray, miArray)

	// Es un slice -> tamaño dinámico
	miSlice := []int{20, 30, 5, 2, 1, 4, 10, 23, 45, 67}
	fmt.Printf("%v del tipo %T\n", miSlice, miSlice)

	// Se puede agregar elementos
	miSlice = append(miSlice, 100)
	fmt.Println("Agregado 100: ", miSlice)
	
	// Remover es mas complicado ya que no existe una funcion que lo haga!!!
	removeSlice(miSlice, 3)
	fmt.Println(miSlice[3:])
	fmt.Println(miSlice[:3])
	fmt.Println("Remover elemento del indice 3 (el valor 2): ", miSlice)
}

// La idea detras de esta funcion es recortar en dos piezas el slice:
// antes del indice y despues del indice. Luego retorna la union de ambas piezas.
func removeSlice(slice [] int, indice int) []int {
	return append(slice[:indice], slice[indice + 1:]...)
}