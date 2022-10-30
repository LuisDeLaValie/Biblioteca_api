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

	colecciones, err := m.ListarColeccion()

	if err != nil {
		cerror := ErrorRes{Error: "Error obteniendo los datos", Cuerpo: err, Mensaje: err.Error()}
		json.NewEncoder(w).Encode(cerror)
	}
	json.NewEncoder(w).Encode(colecciones)
}
func VerColeccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])
	if err == nil {
		coleccion, err := m.VerColeccion(id)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		} else {
			json.NewEncoder(w).Encode(coleccion)
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
		nuevaColeccion, err := m.CrearColeccion(coleccion)

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
			coleccion, err := m.EditarColeccion(id, update)

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
		err = m.EliminarColeccion(id)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}
	}

}
