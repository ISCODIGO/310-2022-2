package main

import "fmt"

func a() {
	fmt.Println("inicia a()")
	b()
	fmt.Println("termina a()")
}

func b() {
	fmt.Println("inicia b()")
	c()
	fmt.Println("termina b()")
}

func c() {
	fmt.Println("inicia c()")
	fmt.Println("termina c()")
}

func main() {
	fmt.Println("inicia main()")
	a()
	fmt.Println("termina main()")
}