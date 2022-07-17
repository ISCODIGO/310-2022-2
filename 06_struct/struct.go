package main

import (
	"fmt"
	"math"
)

type Punto struct {
	x int
	y int
} 

// Parametros de entrada => (x, y int)
// Parametro de salida => (p *Punto) <--- esta es una salida nombrada
func NewPunto(x, y int) (p *Punto) {
	p = &Punto{
		x: x,
		y: y,		
	}

	return
}
//------------------------------------------------------------------------------

type Circulo struct {
	origen Punto
	radio int
}

// Crear un objeto de tipo Circulo
func NewCirculo(p Punto, r int) *Circulo {
	return &Circulo{
		origen: p,
		radio: r,
	}
}


// Metodo para calcular el area
func area_funcion(c Circulo) float64 {
	var area float64
	area = math.Pi * math.Pow(float64(c.radio), 2)
	return area
}

func (c Circulo) area_metodo() float64 {
	var area float64
	area = math.Pi * math.Pow(float64(c.radio), 2)
	return area
}

func main() {
	p := NewPunto(10, 20)
	c := NewCirculo(*p, 40)
	
	fmt.Println(*p)
	fmt.Println(*c)
	fmt.Println("El area del circulo es:", area_funcion(*c), "u^2")
	fmt.Println("El area del circulo es:", c.area_metodo(), "u^2")
}