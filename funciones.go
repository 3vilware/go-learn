package main

import (
	"errors"
	"fmt"
)

//Las funciones sirven igual que en C practicamente

func suma(a, b int16) int16 { //tipo de dato de argumento y tipo de dato de funcion

	return a + b
}

//Funciones que devuelven multiples valores
func maxmin(numeros []int) (int, int) { //tipo de dato que va a retornar cada uno
	var max, min int
	for i, v := range numeros {
		if i == 0 {
			max = v
			min = v
		}
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max, min
}

//Funciones variaticas
func insultos(insulto ...string) {
	fmt.Println(insulto) //por defecto siempre recive un slice y pueden ser los parámetros que sean de los td especificados
}

//Errores en funciones
func division(dividendo, divisor float32) (resultado float32, err error) {
	if divisor == 0 {
		err = errors.New("No se puede dividir por 0")
	}
	resultado = dividendo / divisor

	return
}

//Funciones anonimas multivariable secuenciales :v
func add() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	var n1, n2 int16
	n1 = 30
	n2 = 60
	r := suma(n1, n2)

	fmt.Println("Resultado: ", r)

	n := []int{1, 2, 3, 4, 25, 57, 20}

	max, min := maxmin(n)
	fmt.Println("máximo: ", max)
	fmt.Println("minimo", min)

	insultos("Pendejo", "puto", "koreano", "industrial")

	result, e := division(30, 0)

	if e != nil {
		fmt.Print("Error: ", e)
	} else {
		fmt.Println("Resultado = ", result)
	}

	//Funciones anonimas
	variable := func() {
		fmt.Println("\nsoy anonimo, estoy bien loko!!")
	}

	variable()

	secuencial := add()
	fmt.Println(secuencial())
	fmt.Println(secuencial())
	fmt.Println(secuencial())
}
