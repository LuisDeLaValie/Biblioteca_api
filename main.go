package main

import (
	"fmt"
	"log"
	"net/http"

	"libreria/src/handler"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	// Lirbos
	r.HandleFunc("/libros", handler.ListarLibros).Methods(http.MethodGet)
	r.HandleFunc("/libros/{key:[a-zA-Z0-9]+}", handler.VerLibro).Methods(http.MethodGet)
	r.HandleFunc("/libros", handler.CrearLibro).Methods(http.MethodPost)
	r.HandleFunc("/libros/{key:[a-zA-Z0-9]+}", handler.ActualizarLibro).Methods(http.MethodPut)
	r.HandleFunc("/libros/{key:[a-zA-Z0-9]+}", handler.EliminarLibro).Methods(http.MethodDelete)

	r.HandleFunc("/coleccion", handler.ListarColecciones).Methods(http.MethodGet)
	r.HandleFunc("/coleccion/{key:[a-zA-Z0-9]+}", handler.VerColeccion).Methods(http.MethodGet)
	r.HandleFunc("/coleccion", handler.CrearColeccion).Methods(http.MethodPost)
	r.HandleFunc("/coleccion/{key:[a-zA-Z0-9]+}", handler.ActualizarColeccion).Methods(http.MethodPut)
	r.HandleFunc("/coleccion/{key:[a-zA-Z0-9]+}", handler.EliminarColeccion).Methods(http.MethodDelete)

	r.HandleFunc("/autor", handler.ListarAutores).Methods(http.MethodGet)
	r.HandleFunc("/autor/{key:[a-zA-Z0-9]+}", handler.VerAutor).Methods(http.MethodGet)
	r.HandleFunc("/autor", handler.CrearAutor).Methods(http.MethodPost)
	r.HandleFunc("/autor/{key:[a-zA-Z0-9]+}", handler.ActualizarAutor).Methods(http.MethodPut)
	r.HandleFunc("/autor/{key:[a-zA-Z0-9]+}", handler.EliminarAutor).Methods(http.MethodDelete)

	// http.Handle("/", r)
	fmt.Println("Servidor en Linea")
	log.Fatal(http.ListenAndServe(":8000", r))
}
