//GORM es un ORM para conectarse a bases de datos
// http://gorm.io/docs
package main 

import(
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //caracter _ para que solo se ejecute la funcion init
)

//Declaramos modelos en forma de estructura (por gorm)
// En mayuscula
type Producto struct{
	gorm.Model // ID, CreateAt, UpdateAt, DeleteAt
	CodigoBarras string
	Nombre string
	Precio float32
}

func main(){
	db, err := gorm.Open("mysql","root:ricardo@/godb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Error en la conexión a la base de datos.")
	}
	fmt.Println("Se conecto a la base de datos")

	
	//db.CreateTable(&Producto{}) //creamos la tabla pasando puntero //Ya se creo
	//fmt.Println("Se creo una nueva tabla")

	fmt.Printf("Creando nuevo registro....")
	//db.Create(&Producto{CodigoBarras:"65456", Nombre:"Jabon", Precio:10})
	fmt.Printf(" [OK]\n")

	var p Producto
	db.First(&p) //Ya sabe que es tipo producto y se consulta la tabla producto
	fmt.Println("Datos consultados: ", p)

	p.Precio = 13.5
	db.Save(&p)



	defer db.Close()
}