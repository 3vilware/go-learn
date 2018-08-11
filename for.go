package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		//fmt.Print(i)
	}

	matriz := [3][3]int{}

	var cont int
	cont = 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matriz[i][j] = cont
			cont++
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Print("\n")
	}

	fmt.Println()
	//for continuo (tipo un while)
	var a int
	a = 5

	for a > 0 {
		fmt.Println(a)
		a--
	}

	//For 'forever'
	//se usa para cuando tenemos que estar a la escucha de eventos o conexiones permanentemente

	for {
		fmt.Println("chambea")
	}

	//For in range
	numeros := map[int]string{
		1: "uno",
		2: "dos",
		3: "tres",
	}

	for i, j := range numeros {
		fmt.Println(i, j)
	}
}
