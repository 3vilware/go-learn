package main

import (
	"fmt"

	"github.com/pro/paquetes/funciones" //debe estar dentro de otra carpeta sub siguiente del main
)

func main() {
	computo.PiValue = 3.1416
	fmt.Println(computo.Area(3, computo.PiValue))
}
