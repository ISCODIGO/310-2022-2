package main

import "fmt"

var i int = 0

func main() {
	if i <= 10 {
		i++
		main()
	}
	fmt.Println("Llamada de funcion main")
}