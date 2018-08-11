package main

import "fmt"

func main() {
	//manera para declarar 1
	var numeros [2]int
	numeros[0] = 1
	numeros[1] = 2

	//otra manera
	cadenas := [2]string{"uno", "dos"}

	//Slice (un apuntador a un array (lo hace dinamico))
	// var animales []string //primera forma
	animales := make([]string, 0) //forma más comun para crear slice's make(tipo de dato, capacidad inicial)
	animales = append(animales, "prro")
	animales = append(animales, "chango")

	fmt.Println(cadenas)
	fmt.Println("tamaño: ", len(cadenas))
	fmt.Println("Slice: ", animales)

}
