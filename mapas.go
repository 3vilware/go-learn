package main

import "fmt"

func main() {
	idiomas := make(map[string]string)
	idiomas["es"] = "Español"
	idiomas["en"] = "English"
	idiomas["fr"] = "French"

	numeros := map[int]string{
		1: "uno",
		2: "dos",
		3: "tres",
	}

	fmt.Println(idiomas)

	if idioma, ok := idiomas["es"]; ok { // si ok es if true
		fmt.Println("Existe ", idioma, ok)
	} //devuelve 2 valores

	if numero, ok := numeros[2]; ok {
		fmt.Println("Numero es ", numero)
	}
}
