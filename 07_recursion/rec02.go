package main

import "fmt"

func SumarNBucle(n int) int {
	suma := 0
	for i := 1; i <= n; i++ {
		suma += i
	}

	return suma
}

var recursiones int = 0
func SumarNRecursion(n int) int {
	if n == 1 {
		// Caso base
		recursiones++
		return 1
	}

	// Caso recursivo
	recursiones++
	return n + SumarNRecursion(n - 1)
}


// Suma(4) = 4 + 3 + 2 + 1		=> 4 + Suma(3)
// Suma(3) = 3 + 2 + 1			=> 3 + Suma(2)
// Suma(2) = 2 + 1				=> 2 + Suma(1)
// Suma(1) = 1



func main() {
	N := 100
	x := SumarNBucle(N)
	fmt.Println(x)

	y := SumarNRecursion(N)
	fmt.Println(y)

	fmt.Println("Recursiones: ", recursiones)
}