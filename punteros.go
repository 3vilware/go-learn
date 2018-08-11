package main

import "fmt"

func main() {
	a := 3
	b := 5
	fmt.Println("Antes de duplicar a vale ", a)
	duplicar(a)
	fmt.Println("Despues de duplicar, a vale ", a)

	fmt.Println("b antes de duplicar con puntero vale ", b)
	duplicarPuntero(&b)
	fmt.Println("b después de duplicar con puntero vale", b)
}

func duplicar(x int) {
	x = x * 2
	fmt.Println("valor duplicado = ", x)
}

func duplicarPuntero(x *int) {
	*x = *x * 2
	fmt.Println("valor duplicado = ", *x)
}
