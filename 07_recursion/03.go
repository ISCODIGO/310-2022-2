/*
sumar(1): 										1 ==> 1 (caso base)
sumar(7):               7 + 6 + 5 + 4 + 3 + 2 + 1
sumar(8):           8 + 7 + 6 + 5 + 4 + 3 + 2 + 1
sumar(9):       9 + 8 + 7 + 6 + 5 + 4 + 3 + 2 + 1 ==> 9 + sumar(8)
sumar(10): 10 + 9 + 8 + 7 + 6 + 5 + 4 + 3 + 2 + 1 ==> 10 + sumar(9)
*/

package main

import "fmt"

func main() {
	
	fmt.Println(sumar_rec(10))
	fmt.Println(sumar_bucle(10))
}

// Es T: O(n) / S: O(n)
func sumar_rec(n int) (int) {
	if n == 1 {
		return n
	} else {
		return n + sumar_rec(n - 1)
	}
}

// Es T: O(n) / S: O(1)
func sumar_bucle(n int) (int) {
	suma := 0
	for i := 1; i <= n; i++ {
		suma += i
	}

	return suma
}