package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct{
	message string
}

func(m messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, m.message)
}

func main() {
	mux := http.NewServeMux()
	m1 := &messageHandler{message: "<h1>hola desde el handler</h1>"}
	m2 := &messageHandler{message: "<h1>estamos en otra ruta</h1>"}
	
	mux.Handle("/saludar", m1) // M1 es el contenido de lo que quiero que renderize
	mux.Handle("/otra", m2) // le paso otra ruta y contenido
	
	log.Println("Corriendo servidor... [OK]")
	log.Println(http.ListenAndServe(":8080", mux)) // le paso el multiplexor
}
