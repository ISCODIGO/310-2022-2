package array

import "testing"

func TestListaVacia(t *testing.T) {
	lista := NewListaArray()

	// Tendra un size == 0
	if lista.Length() != 0 {
		t.Error("Se esperaba un size cero")
	}

}

func TestAppend(t *testing.T) {
	lista := NewListaArray()

	lista.Append(10)
	lista.Append(20)

	// Tendra 2 elementos
	if lista.Length() != 2 {
		t.Error("Se esperaban 2 elementos")
	}

	lista.MoveToPos(0)
	valor, _ := lista.GetValue()
	if valor != 10 {
		t.Error("Se esperaba que el primero fuese 10")
	}

	lista.MoveToPos(1)
	valor, _ = lista.GetValue()
	if valor != 20 {
		t.Error("Se esperaba que el segundo fuese 20")
	}
}

func TestRemove(t *testing.T) {
	lista := NewListaArray()

	lista.Append(10)
	lista.Append(20)

	lista.MoveToStart()
	removido, error := lista.Remove()  // Remueve el primero
	
	if error != nil {
		t.Error("Este error no debio ocurrir:", error)
	}
	
	if removido != 10 {
		t.Error("Debio removerse el 10 pero se removio el", removido)
	}

	if lista.Length() != 1 {
		t.Error("Debia quedar solo 1 elemento y hay", lista.Length())
	}

	removido_2, error := lista.Remove()

	if error != nil {
		t.Error("Este error no debio ocurrir:", error)
	}

	if removido_2 != 20 {
		t.Error("Debio removerse el 20 pero fue", removido_2)
	} 

	if lista.Length() != 0 {
		t.Error("Debia quedar vacio")
	}

	_, error = lista.Remove()
	
	if error == nil {
		t.Error("Debio generarse un error")
	}
}

func TestClear(t *testing.T) {
	lista := NewListaArray()

	lista.Append(10)
	lista.Append(20)
	lista.Clear()

	if lista.Length() == 0 {
		t.Error("Debia quedar vacio")
	}
}

func TestMoveNext(t *testing.T) {
	lista := NewListaArray()

	lista.Append(10)
	lista.Append(20)
	lista.Append(30)

	lista.MoveToStart()
	lista.Next()
	actual, _ := lista.GetValue()
	if actual != 20 {
		t.Error("Next: El valor debio ser 20")
	}

	lista.Next()
	actual, _ = lista.GetValue()
	if actual != 30 {
		t.Error("Next: El valor debio ser 30")
	}

	lista.Next()
	actual, _ = lista.GetValue()
	if actual != 30 {
		t.Error("Next: no debe moverse al llegar al final. Sale", actual)
	}
}

func TestMovePrev(t *testing.T) {

	lista := NewListaArray()

	lista.Append(10)
	lista.Append(20)
	lista.Append(30)
	lista.MoveToEnd()

	actual, _ := lista.GetValue()
	if actual != 30 {
		t.Error("MoveToEnd: el valor debio ser 30 pero aparece", actual)
	}

	lista.Prev()
	actual, _ = lista.GetValue()
	if actual != 20 {
		t.Error("Prev: el valor debio ser 20 pero aparece", actual)
	}

	lista.Prev()
	actual, _ = lista.GetValue()
	if actual != 10 {
		t.Error("Prev: el valor debio ser 10 pero aparece", actual)
	}

	// Ese es el comportamiento del libro: al llegar al principio no retrocede mas
	lista.Prev()
	actual, _ = lista.GetValue()
	if actual != 10 {
		t.Error("Prev: No debe moverse al llegar al principio")
	}
}

func TestInsert(t *testing.T) {
	lista := NewListaArray()
	lista.Append(10)
	lista.Append(30)

	lista.MoveToPos(1)
	lista.Insert(20)

	valor, _ := lista.GetValue()
	if lista.Length() != 3 {
		t.Error("Deben tener 3 elemento pero hay", lista.Length())
	}

	if valor != 20 {
		t.Error("En la posicion 1 debe estar el 20 pero esta", valor)
	}
}