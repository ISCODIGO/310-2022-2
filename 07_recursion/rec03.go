package main

import "fmt"

var recursiones int = 0

func fibonacci(n int) int {
	// Caso base
	if n < 2 {
		recursiones++
		return n
	}

	recursiones++

	// Caso recursivo
	return fibonacci(n - 1) + fibonacci(n - 2)
}

func main() {
	fmt.Println("Escriba un valor para N")

	var numero int
	fmt.Scan(&numero)

	fibo := fibonacci(numero)
	fmt.Printf("Fibonacci (%v) = %v\n", numero, fibo)

	fmt.Println("Recursiones generadas:", recursiones)
}