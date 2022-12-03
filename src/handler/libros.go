package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	m "libreria/src/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListarLibros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		if r := recover(); r != nil {
			json.NewEncoder(w).Encode(r)
		}
	}()

	var l m.Libro
	id := r.URL.Query().Get("search")
	_only := r.URL.Query().Get("only")
	only, _ := strconv.ParseBool(_only)

	libros := l.Listar(id, only)
	json.NewEncoder(w).Encode(libros)
}
func VerLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer func() {
		if r := recover(); r != nil {
			json.NewEncoder(w).Encode(r)
		}
	}()

	var libro m.Libro
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])
	if err != nil {
		panic(m.ErrorRes{Error: "key no valida", Cuerpo: err, Mensaje: err.Error()})
	}

	libro.Ver(id)
	json.NewEncoder(w).Encode(libro)
}
func CrearLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		if r := recover(); r != nil {
			json.NewEncoder(w).Encode(r)
		}
	}()

	// Obtener datos del formulario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		// Preparar formulario y mandar la informacion
		var libro m.LibroFormulario
		json.Unmarshal(reqBody, &libro)
		nuevoLibro := libro.Crear()
		json.NewEncoder(w).Encode(nuevoLibro)

	}
}
func ActualizarLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		if r := recover(); r != nil {
			json.NewEncoder(w).Encode(r)
		}
	}()

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])

	if err == nil {
		reqBody, err := ioutil.ReadAll(r.Body)

		if err == nil {
			var update m.LibroFormulario
			json.Unmarshal(reqBody, &update)
			libro := update.Editar(id)
			json.NewEncoder(w).Encode(libro)
		} else {
			json.NewEncoder(w).Encode(err.Error())
		}
	} else {
		json.NewEncoder(w).Encode(err.Error())
	}
}
func EliminarLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		if r := recover(); r != nil {
			json.NewEncoder(w).Encode(r)
		}
	}()

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])
	if err == nil {
		var l m.Libro
		err = l.Eliminar(id)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}
	}
}
