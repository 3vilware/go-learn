package computo //el nombre con el que se invoca en el main

var PiValue float32

func Area(radio, pi float32) float32 {
	a := (pi * (radio * radio)) / 2

	return a
}
