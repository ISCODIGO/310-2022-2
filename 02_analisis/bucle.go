package main

import "fmt"

func main() {
	iteracion := 0
	const N = 100000

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			iteracion += 1
		}
	}

	fmt.Printf("f(%v) = %v", N, iteracion)
}

/*

	for i := 1; i < 2N; i++ {
		iteracion += 1
	}

	n		f(n)
	100		200
	500		1000
	m		2*m

	f(n) ~ 6n + 5 es O(n)
--------------------------------

	for i := 1; i < N; i *= 2 {
		iteracion += 1
	}

	n		 	f(n)
	100			7
	500			9
	10000		14

	f(n) ~ log₂(n) -> O(log₂(n))

	--------------------------------

	for i := N; i >= 1; i /= 2 {
		iteracion += 1
	}

	n		 	f(n)
	100			7
	500			9
	10000		14

	f(n) ~ log₂(n) -> O(log₂(n))

	--------------------------------

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			iteracion += 1
		}
	}

	n		 	f(n)
	10			100
	100			10000
	500			250000
	10000		100,000,000
	n			n*n = n^2


	f(n) ~ n^2 => O(n^2)

	-----------------------

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				iteracion += 1
			}
		}
	}

	n		 	f(n)
	2			8
	10			1000
	100			1000000
*/