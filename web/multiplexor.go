package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	log.Println("Corriendo servidor... [OK]")
	log.Println(http.ListenAndServe(":8080", mux)) // le paso el multiplexor

}
