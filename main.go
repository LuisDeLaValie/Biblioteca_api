package main

import (
	"net/http"
	"log"
	"fmt"

	"libreria/src/handler"

	"github.com/gorilla/mux"

)

func main(){

	r := mux.NewRouter()

    r.HandleFunc("/libros", handler.ListarLibros).Methods(http.MethodGet)
    r.HandleFunc("/libros/{key:[a-zA-Z0-9]+}", handler.VerLibro).Methods(http.MethodGet)
    r.HandleFunc("/libros", handler.CrearLibro).Methods(http.MethodPost)
    r.HandleFunc("/libros", handler.ActualizarLibro).Methods(http.MethodPut)
    r.HandleFunc("/libros", handler.EliminarLibro).Methods(http.MethodDelete)
    
	// http.Handle("/", r)
	fmt.Println("Servidor en Linea")
	log.Fatal(http.ListenAndServe(":8000",r))
}