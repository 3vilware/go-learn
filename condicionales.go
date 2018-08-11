package main

import (
	"fmt"
)

func main() {
	//:= es asignale el tipo de dato que le corresponda a la variable

	isValid := true
	var a uint8
	a = 1

	if !isValid {
		fmt.Println("Es valido viejon")
	} else { // el else siempre se tiene que poner asi
		fmt.Println("se fue al else!")
	}

	switch a {
	case 1:
		fmt.Printf("valor es uno %d", a)
	case 2:
		fmt.Println("Nel la neta no")
	default:
		fmt.Println("al default")
	}
}
