/* 
Mostrar la existencia de las variables globales dentro del contexto de un 
archivo de Go.
*/
package main

import "fmt"

// variable global
var a int = 10

func main() {
	otra()
	fmt.Println("Llamado desde main de a: ", a)
}

func otra() {
	a = 20
	fmt.Println("Llamado desde otra() de a: ", a)
}