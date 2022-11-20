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

func ListarColecciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var col m.Coleccion
	colecciones, err := col.Listar()

	if err != nil {
		cerror := m.ErrorRes{Error: "Error obteniendo los datos", Cuerpo: err, Mensaje: err.Error()}
		json.NewEncoder(w).Encode(cerror)
	}
	json.NewEncoder(w).Encode(colecciones)
}
func VerColeccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])
	if err == nil {
		var col m.Coleccion
		err := col.Ver(id)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		} else {
			json.NewEncoder(w).Encode(col)
		}
	} else {
		json.NewEncoder(w).Encode(err.Error())
	}

}
func CrearColeccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Obtener datos del formulario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		// Preparar formulario y mandar la informacion
		var coleccion m.ColeccionFormulario
		json.Unmarshal(reqBody, &coleccion)
		nuevaColeccion, err := coleccion.Crear()

		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		} else {
			json.NewEncoder(w).Encode(nuevaColeccion)
		}

	}

}
func ActualizarColeccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])

	if err == nil {
		reqBody, err := ioutil.ReadAll(r.Body)

		if err == nil {
			var update m.ColeccionFormulario
			json.Unmarshal(reqBody, &update)
			coleccion, err := update.Editar(id)

			if err != nil {
				json.NewEncoder(w).Encode(err.Error())
			} else {
				json.NewEncoder(w).Encode(coleccion)
			}
		} else {
			json.NewEncoder(w).Encode(err.Error())
		}
	} else {
		json.NewEncoder(w).Encode(err.Error())
	}

}
func EliminarColeccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])
	if err == nil {
		var col m.Coleccion
		err = col.Eliminar(id)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}
	}

}
