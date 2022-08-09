package main

import "fmt"

var recursiones int = 0

func fibonacci(n int, buffer []int) int {
	// O(n)
	if n < 2 {
		recursiones++
		return n
	}

	recursiones++

	// Se crea el buffer
	if len(buffer) == 0 {
		buffer = make([]int, n + 1)  // el slice se rellena de ceros
		buffer[0] = 0
		buffer[1] = 1
	}

	if buffer[n] == 0 {
		buffer[n] = fibonacci(n - 1, buffer) + fibonacci(n - 2, buffer)
	}

	return buffer[n]
}

func main() {
	var numero int
	
	fmt.Println("Escriba un valor para N")
	fmt.Scan(&numero)
	var memoria []int
	fibo := fibonacci(numero, memoria)
	fmt.Printf("Fibonacci (%v) = %v\n", numero, fibo)
	fmt.Println("Recursiones generadas:", recursiones)
}