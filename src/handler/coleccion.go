package handler

import (
	"encoding/json"
	"io/ioutil"
	"libreria/src/model"
	"libreria/src/model/coleccion"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/gorilla/mux"
)

func ListarColecciones(w http.ResponseWriter, r *http.Request) {
	Headers(&w)
	var col coleccion.Coleccion
	search := r.URL.Query().Get("search")

	colecciones, err := col.Listar(search)

	if err != nil {
		cerror := model.ErrorRes{Titulo: "Error obteniendo los datos", Cuerpo: err, Mensaje: err.Error()}
		json.NewEncoder(w).Encode(cerror)
	} else {
		res := struct {
			Total     int                     `json:"total"`
			Coleccion coleccion.ListColeccion `json:"coleccion"`
		}{
			Total:     len(colecciones),
			Coleccion: colecciones,
		}
		json.NewEncoder(w).Encode(res)
	}
}
func VerColeccion(w http.ResponseWriter, r *http.Request) {
	Headers(&w)
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])
	if err == nil {
		var col coleccion.Coleccion
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
	Headers(&w)

	// Obtener datos del formulario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		// Preparar formulario y mandar la informacion
		var coleccion coleccion.ColeccionFormulario
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
	Headers(&w)

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])

	if err == nil {
		reqBody, err := ioutil.ReadAll(r.Body)

		if err == nil {
			var update coleccion.ColeccionFormulario
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
	Headers(&w)

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])
	if err == nil {
		var col coleccion.Coleccion
		all, _ := strconv.ParseBool((vars["all"]))
		err = col.Eliminar(id, all)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}
	}

}
