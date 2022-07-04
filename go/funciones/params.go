package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v %T\n", funcion1(4), funcion1(4))

	var x = 10
	var y = 20
	
	var v1 int
	var v2 int

	v1, v2 = function2(x, y)
	fmt.Println(v1, v2)

	hola := "hola"
	adios := "adios"
	
	// esta es la ejecucion de un paso por referencia
	swap(&hola, &adios) 
	fmt.Println(hola, adios)

	// funciones variadicas
	fmt.Println(sumar(1))
	fmt.Println(sumar(1, 2))
	fmt.Println(sumar(1, 2, 3))
	var arr []int = []int{1, 2, 3, 4}
	fmt.Println(sumar(arr...))

	// funciones anonimas, se puede pasar como variable
	duplicar := func (a int ) int {
		return 2 * a
	}

	fmt.Println(duplicar(10))
}


// entrada: int y salida: float32. Las funciones pueden tener como opcionales
// tanto sus entradas como salidas, pero una vez se indica la salida es necesario
// colocar un return
func funcion1(entrada int) float32 {
	return float32(entrada)
}

// como los retornos llevan nombre no es necesario definirlos en el return
// pero es necesario utilizarlos en el cuerpo de la funcion
func function2(arg1 int, arg2 int) (a int, b int) {
	a = arg1 + arg2
	b = arg1 - arg2

	return
}

// Paso por referencia de una funcion. Se pasan dos punteros string
func swap(a, b *string) {
	temp := *a
	*a = *b
	*b = temp
}

// Funcion que permite varios parametros (variadicas). Los 3 puntos se denominan
// elipsis
func sumar(numeros ...int) int {
	suma := 0
	for _, valor := range numeros {
		suma += valor
	}
	return suma
}
