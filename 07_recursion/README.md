# Funciones Recursivas

## `func.go`
Este archivo demuestra que internamente el llamado de las funciones ocurre por medio de una pila que almacena las funciones en espera.

## `rec00.go`
Cuando una función se llama a sí misma. Pero no tiene caso base por lo que se genera una recursión infinita.

## `rec01.go`
En este caso se genera un caso base. Lo importante es que la recursión debe llevarnos al caso base.

```
recursión -> caso base
```

Algo de hacer notar es que el incremento de `i` debe hacerse antes del llamado recursivo ya que de otra forma causaría recursión infinita.

## `rec02.go`
En este archivo se establece una comparacion entre una solución iterativa y una recursiva para resolver la suma de enteros que van desde 1 hasta N.


## `rec03.go`
En este archivo tenemos el Fibonacci de toda la vida: Una serie que genera números de la siguiente serie:
```
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
```
Esta solución es simple en términos de código es similar a su representación matemática. Pero la eficiencia de esta solución es O(2<sup>n</sup>). Se puede notar los tiempos altos en la solución para valores altos de `n`.

## `rec04.go`
Esta es una alternativa para generar números de la serie de Fibonacci. Utiliza una especie de **buffer** para evitar recalcular valores previos. Esto aumenta la eficiencia a O(n).

## `rec05.go`
En este ejemplo se puede observar el uso de la recursión para problemas donde la iteración puede resultar complicada: resolver de forma ingenua un laberinto.

### Nuevos elementos
- Uso de un slice de 2 dimensiones.
- Para crear un slice se puede utilizar [make](https://go.dev/tour/moretypes/13).