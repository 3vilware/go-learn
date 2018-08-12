//Handlerfunc son un tipo de dato que convierte funciones en manejadores http. Requisito una firma en especifico
package main

import(
	"net/http"
	"fmt"
	"log"
)

func messageHandlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Este es un handler func</h1>")
}

func customMessage(message string) http.Handler{
	return http.HandlerFunc( //para func
		func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintf(w, message)
		},
	)
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", messageHandlerFunc)//tipo para main
	mux.Handle("/custom", customMessage("My niu mensaje"))

	log.Println("Corriendo servidor... [OK]")
	log.Println(http.ListenAndServe(":8080", mux)) // le paso el multiplexor

}