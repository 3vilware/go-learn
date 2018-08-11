package main

import (
	"errors"
	"fmt"
)

func main() {
	e, conectado, db := conectarDB("ruck", "test")

	if conectado {
		fmt.Println("nombre de la base de datos: ", db)
	} else {
		fmt.Println("No se pudo realizar la conexion")
		fmt.Println("Error: ", e)
	}

	// defer ejecuta esa funcion o funciones independientemente si hubo un error que tronara antes. Evita el gasto de recursos
	defer closeDB()
}

func conectarDB(user, pass string) (error, bool, string) {
	var valido bool
	var dbname string
	var err error

	if user == "rick" && pass == "test" {
		valido = true
		dbname = "produccion"
	} else {
		valido = false
		dbname = "Na"
		err = errors.New("Usuario o contraseña incorrectos")
	}
	return err, valido, dbname
}

func closeDB() {
	fmt.Println("Conexión cerrada")
}
