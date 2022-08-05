package main

import "fmt"

const N int = 10

// Es O(n)
func rec2(conteo int) {
	fmt.Println("inicia rec()")

	// caso base
	if conteo < N {
		rec2(conteo + 1)
	}

	fmt.Println("fin rec()")
}


func main() {
	// rec2(0) -> rec2(1) -> rec2(2) -> rec2(3) -> ... -> rec2(N)
	rec2(0)
}