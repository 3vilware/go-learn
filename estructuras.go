package main

import "fmt"

type Persona struct {
	nombre string
	edad   uint8
}

func main() {
	var yo Persona
	el := Persona{nombre: "dante", edad: 6}

	yo.nombre = "Ricardo Amador"
	yo.edad = 21

	fmt.Println(yo)
	fmt.Println(el)

}
