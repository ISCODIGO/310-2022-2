package main

import "fmt"

func main() {
	rec()
	fmt.Println(sumar(10))
	fmt.Println(sumar_bucle(10))
}

func a() {
	fmt.Println("inicia a()")
	b()
	fmt.Println("fin a()")
}

func b() {
	fmt.Println("inicia b()")
	c()
	fmt.Println("fin b()")
}

func c() {
	fmt.Println("inicia c()")
	fmt.Println("fin c()")
}


func rec() {
	fmt.Println("inicia rec()")
	rec()
	fmt.Println("termina rec()")
}

func sumar(n int) (int) {
	if n == 1 {  // caso base
		return 1
	} else {  // caso recursivo
		return n + sumar(n-1)
	}
}

