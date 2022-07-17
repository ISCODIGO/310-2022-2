package main

import "fmt"

func main() {
	var miSlice []int

	fmt.Println(miSlice)
	fmt.Println("Capacidad: ", cap(miSlice), "Tamaño: ", len(miSlice))

	for i := 0; i < 97; i++ {
		// Insertamos elementos en un slice
		miSlice = append(miSlice, i)
		fmt.Println("Capacidad: ", cap(miSlice), "Tamaño: ", len(miSlice))
	}
}