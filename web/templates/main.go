package main

import (
	"html/template"
	"log"
	"net/http"
)

//Estructura de ejemplo
type Persona struct {
	Nombre string
	Edad   uint8
}

func renderizar(w http.ResponseWriter, r *http.Request) {
	p := &Persona{"Ricardo", 21}
	t, err := template.ParseFiles("./views/index.html")

	//Si ocurrio un error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//t.Execute(w, nil) //ejecutamos template
	t.Execute(w, p) //ejecutamos template
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", renderizar)

	log.Println("Corriendo servidor template... [OK]")
	log.Println(http.ListenAndServe(":8080", mux)) // le paso el multiplexor

}
