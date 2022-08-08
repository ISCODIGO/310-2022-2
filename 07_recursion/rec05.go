package main

import (
	"bytes"
	"fmt"
)

const MURO int = 0
const LIBRE int = 1
const PASO int = 2

type Laberinto struct {
	n int
	tablero [][]int
}

//------------------------------------------------------------------------------

// FUNCIONES AUXILIARES

// Funcion auxiliar para rellenar un slice 2D con una dimension N x M
func rellenarSlice2D(n, m int) [][]int {
	matriz := make([][]int, n)
	for i := range matriz {
		matriz[i] = make([]int, m)
	}
	return matriz
}

func StringSlice2D(matriz [][]int) string {
	var b bytes.Buffer

	for _, col := range matriz {
		for _, val := range col {
			b.WriteString(fmt.Sprintf(" %v ", val))
		}
		b.WriteString("\n")
	}

	return b.String()
}

//------------------------------------------------------------------------------

// FUNCIONES DE STRUCT

func NewLaberinto(matriz [][]int) *Laberinto {
	laberinto := new(Laberinto)
	// Se inicializa un slice con un tamaño determinado
	laberinto.tablero = make([][]int, len(matriz))
	// Para copiar un slice a otro se debe tener la misma dimension
	copy(laberinto.tablero, matriz)
	laberinto.n = len(matriz)
	
	return laberinto
}

func (lb Laberinto) String() string {
	return StringSlice2D(lb.tablero)
}

func (lb Laberinto) esPermitido (x, y int) bool {
	// Esta funcion permite que saber si los los ejes estan dentro de los limites
	// y que la celda del tablero esta libre
	if (x >= 0 && x < lb.n) && (y >= 0 && y < lb.n && lb.tablero[x][y] == LIBRE) {
		return true
	}

	return false
}

func (lb *Laberinto) Resolver() bool {
	// El arreglo SOLUCION sirve para almacenar el recorrido de la solucion
	// asi el tablero se mantiene inalterable
	solucion := make([][]int, lb.n)
	copy(solucion, lb.tablero)

	if !lb.resolverRecursivo(0, 0, solucion) {
		return false
	}

	fmt.Println(StringSlice2D(solucion))
	return true
}

func (lb *Laberinto) resolverRecursivo(x int, y int, sol [][]int) bool {
	// Caso base: llegar a la meta
	if (x == lb.n - 1 && y == lb.n - 1 && lb.tablero[x][y] == LIBRE) {
		sol[x][y] = PASO
		return true
	}

	if lb.esPermitido(x, y) {
		sol[x][y] = PASO

		// Si avanza por 'X' (hacia la derecha)
		if lb.resolverRecursivo(x + 1, y, sol) {
			return true
		}

		// Si avanza por 'Y' (hacia abajo)
		if lb.resolverRecursivo(x, y + 1, sol) {
			return true
		}

		// Si el camino esta cerrado dar marcha atras (backtracking)
		sol[x][y] = lb.tablero[x][y]
		return false
	}

	return false
}

func main() {
	arr := [][]int {
		{ 1, 1, 1, 1, 1, 0, 1, 0 }, 
		{ 1, 0, 0, 0, 0, 0, 1, 0 }, 
		{ 1, 1, 1, 1, 1, 1, 1, 0 }, 
		{ 1, 1, 1, 1, 1, 1, 1, 0 },
		{ 0, 0, 1, 1, 0, 0, 0, 0 },
		{ 0, 0, 1, 1, 1, 0, 0, 0 },
		{ 1, 1, 1, 1, 1, 1, 0, 0 },
		{ 1, 0, 0, 0, 1, 1, 1, 1 },
	}

	fmt.Printf("ESTADO INICIAL... \n([%v] -> espacio disponible y [%v] -> paredes)\n", LIBRE, MURO)
	fmt.Println(StringSlice2D(arr))

	fmt.Printf("SOLUCION.. \n([%v] -> camino recorrido)\n", PASO)
	laberinto := NewLaberinto(arr)
	laberinto.Resolver()
}