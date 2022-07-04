package main

import "fmt"

func main() {
	// Es un array -> tamaño fijo
	numeros := [...]int{20, 30, 5, 2, 1, 4, 10, 23, 45, 67}
	fmt.Println("Escriba un numero entero...")
	var numero int
	fmt.Scan(& numero)
	fmt.Println("Se leyo el valor", numero)

	var fue_hallado bool  // false, por defecto
	var pasos int  // 0, por defecto
	
	// foreach
	for i, valor := range numeros {
		pasos += 1
		if valor == numero {
			fmt.Println("Se encontro en la posicion:", i)
			fue_hallado = true
			break
		}
	}

	if !fue_hallado {
		fmt.Println("No se encontro el numero: ", numero)
	}
	fmt.Printf("Se hicieron %v pasos", pasos)
}