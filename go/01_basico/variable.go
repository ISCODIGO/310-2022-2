package main

import "fmt"

var globalvar = 40.9
const globalconst = 222

func main() {
	var var1 int8 = 15
	var bval bool
	var fval float32
	var sval string

	edad := 24
	altura := 1.7
	nombre := "Mario"

	fmt.Printf("var1 ---> %v(%T)\n", var1, var1)
	fmt.Printf("%v(%T)\n", edad, edad)
	fmt.Printf("%v(%T)\n", altura, altura)
	fmt.Printf("%v(%T)\n", bval, bval)
	fmt.Printf("%v(%T)\n", fval, fval)
	fmt.Printf("'%v'(%T)\n", sval, sval)
	fmt.Printf("'%v'(%T)\n", nombre, nombre)
	fmt.Printf("'%v'(%T)\n", globalvar, globalvar)
	fmt.Printf("'%v'(%T)\n", globalconst, globalconst)
	//fmt.Printf("Suma es: %v", suma(4.9, 3))
	fmt.Printf("Resta es: %v\n", resta(0, 257))
	
	test()
	fmt.Println("Despues de ejecutado test() el valor de globalvar es ", globalvar)
	var s1 = "hola"
	var s2 = "adios"
	swap(s1, s2)
	fmt.Printf("\nSwap-%v - %v", s1, s2)
}

func test() {
	globalvar = 20
	//globalconst = 7
}

// [func] [nombre-funcion] ([arg tipo], [arg2 tipo]) [tipo retorno] {[cuerpo]}
func suma(a int, b int) int {
	return a + b
}

func resta(a, b int) int {
	return a - b
}

func swap(a, b string) (x, y string) {
	return b, a
}