/* 

Serie Fibonacci: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...

Caso base
F(0) = 0
F(1) = 1
---------------------------
Caso recursivo
F(2) = 1 -> 0 + 1 -> F(0) + F(1)
F(3) = 2 -> 1 + 1 -> F(1) + F(2)
F(4) = 3 -> 2 + 1 -> F(2) + F(3)
F(5) = 5 -> 3 + 2 -> F(3) + F(4)
F(6) = 8 -> 5 + 3 -> F(4) + F(5)

F(n) = F(n - 1) + F(n - 2)

-----------------------------


N		recursiones
1		1
5		15
10		177
20		6,765
30		2,692,527
40		331,160,281

*/

package main

import "fmt"

var recursiones int = 0

func fibonacci(n int) int {
	// O(2^n)

	
	// Caso base
	if n < 2 {
		recursiones++
		return n
	}

	// Caso recursivo
	recursiones++
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