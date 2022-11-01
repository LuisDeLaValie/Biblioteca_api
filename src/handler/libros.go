package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	m "libreria/src/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/gorilla/mux"
)

type ErrorRes struct {
	Error   string `json:"error"`
	Mensaje string `json:"mensaje,omitempty"`
	Cuerpo  error  `json:"cuerpo,omitempty"`
}

func ListarLibros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var l m.Libro
	libros, err := l.Listar()

	if err != nil {
		cerror := ErrorRes{Error: "Error obteniendo los datos", Cuerpo: err, Mensaje: err.Error()}
		json.NewEncoder(w).Encode(cerror)
	}
	json.NewEncoder(w).Encode(libros)
}
func VerLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])
	if err == nil {
		var libro m.Libro
		err := libro.Ver(id)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		} else {
			json.NewEncoder(w).Encode(libro)
		}
	} else {
		json.NewEncoder(w).Encode(err.Error())
	}
}
func CrearLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Obtener datos del formulario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		// Preparar formulario y mandar la informacion
		var libro m.LibroFormulario
		json.Unmarshal(reqBody, &libro)
		nuevoLibro, err := libro.Crear()

		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		} else {
			json.NewEncoder(w).Encode(nuevoLibro)
		}

	}
}
func ActualizarLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])

	if err == nil {
		reqBody, err := ioutil.ReadAll(r.Body)

		if err == nil {
			var update m.LibroFormulario
			json.Unmarshal(reqBody, &update)
			libro, err := update.Editar(id)

			if err != nil {
				json.NewEncoder(w).Encode(err.Error())
			} else {
				json.NewEncoder(w).Encode(libro)
			}
		} else {
			json.NewEncoder(w).Encode(err.Error())
		}
	} else {
		json.NewEncoder(w).Encode(err.Error())
	}
}
func EliminarLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
