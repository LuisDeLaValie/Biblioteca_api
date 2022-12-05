package handler

import (
	"encoding/json"
	"io/ioutil"
	"libreria/src/model"
	"libreria/src/model/autor"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/gorilla/mux"
)

func ListarAutores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var a autor.Autor

	search := r.URL.Query().Get("search")

	colecciones, err := a.Listar(search)

	if err != nil {
		cerror := model.ErrorRes{Titulo: "Error obteniendo los datos", Cuerpo: err, Mensaje: err.Error()}
		json.NewEncoder(w).Encode(cerror)
	}
	json.NewEncoder(w).Encode(colecciones)
}
func VerAutor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])
	if err == nil {
		var autor autor.Autor
		autor.Ver(id)

		json.NewEncoder(w).Encode(autor)
	} else {
		json.NewEncoder(w).Encode(err.Error())
	}

}
func CrearAutor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Obtener datos del formulario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		// Preparar formulario y mandar la informacion
		var autor autor.Autor
		json.Unmarshal(reqBody, &autor)
		err := autor.Crear()

		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		} else {
			json.NewEncoder(w).Encode(autor)
		}

	}

}
func ActualizarAutor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])

	if err == nil {
		reqBody, err := ioutil.ReadAll(r.Body)

		if err == nil {
			var update autor.Autor
			json.Unmarshal(reqBody, &update)
			err := update.Editar(id)

			if err != nil {
				json.NewEncoder(w).Encode(err.Error())
			} else {
				json.NewEncoder(w).Encode(update)
			}
		} else {
			json.NewEncoder(w).Encode(err.Error())
		}
	} else {
		json.NewEncoder(w).Encode(err.Error())
	}

}
func EliminarAutor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])
	if err == nil {
		var autor autor.Autor
		err = autor.Eliminar(id)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}
	}

}
