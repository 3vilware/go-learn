package main

import "fmt"

type Prro struct {
	nombre  string
	ladrido string
}

func (p Prro) ladrar() {
	fmt.Println("no pos wuau")
}

func (p *Prro) nombrar(name string) { //Pasamos el puntero para que modifique el objeto y no sea solo una copia
	//p.nombre = "Dante"
	p.nombre = name
	fmt.Println("Me llamo ", p.nombre)
}

func main() {
	var perro Prro
	perro.ladrar()
	perro.nombrar("Dante")
	fmt.Println(perro)
}
