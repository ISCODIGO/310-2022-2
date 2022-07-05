/*
Esta es la forma de generar numeros aleatorios en Go. Utilizando una semilla
que varie en cada ejecucion.
*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // [var] [nombre de la variable] [tipo] = [valor]
    var semilla int64 = time.Now().UnixMilli()       
    rand.Seed(semilla)
    fmt.Println("Numero pseudo-aleatorio:", rand.Intn(10))    
}