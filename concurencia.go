package main

import (
	"fmt"
	"time"
)

func main() {
	go numeros() //con go podemos correr un hilo aparte para esa funcion 
	mensaje()
}

func numeros() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func mensaje() {
	var nombre string
	fmt.Printf("Nombre a saludar: ")
	fmt.Scan(&nombre)
	fmt.Println("Salu2", nombre)
}
