package basica

import (
	"testing"
)

func TestSuma(t *testing.T) {
	valor := Sumar(10, 5)
	esperado := 15

	if valor != esperado {
		t.Error("Se esperaba ", esperado, " pero se obtuvo ", )
	}
}