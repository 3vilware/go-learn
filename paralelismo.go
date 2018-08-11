package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4) //Cantidad de nucleos a utilizar para ejecucion
	var wg sync.WaitGroup //cola de rutinas
	numeros := []int{123, 43, 6456, 76, 89, 987, 235, 459325324}
	wg.Add(len(numeros)) //añadimos la cantidad de elementos del grupo de espera

	for _, v := range numeros {
		go func(a int) {
			defer wg.Done() //cuando acabe esta rutina saca 1 de la cola de rutinas a ejecutar
			fmt.Println(suma(a))
		}(v)
	}
	wg.Wait() //Que espere a que todas las rutinas terminen
	fmt.Println("Terminaron todos =D")

}

func suma(a int) int {
	return a + a*2
}
