package main

import (
	"net/http"
	"log"
)

func main(){
	http.Handle("/", http.FileServer(http.Dir("public")))  //ruta de la raiz de archivos
	log.Println("Corriendo servidor... [OK]")	
	log.Println(http.ListenAndServe("localhost:8080",nil)) //que me diga si hay un error //localhost se puede dejar en blanco

}