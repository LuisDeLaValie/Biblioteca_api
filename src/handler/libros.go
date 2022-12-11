package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"libreria/src/model"
	"libreria/src/model/libro"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListarLibros(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListarLibros")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	Headers(&w)
	w.Header().Set("Accept", "*/*")

	var l libro.Libro
	_only := r.URL.Query().Get("only")
	if _only == "" {
		_only = "0"
	}
	if only, err := strconv.ParseBool(_only); err != nil {
		elerro := model.ErrorRes{
			Status: http.StatusBadRequest,
			Titulo: "dato envalido ",
			Cuerpo: err,
		}
		elerro.Response(w)
	} else {
		id := r.URL.Query().Get("search")
		if libros, err := l.Listar(id, only); err != nil {
			elerro := model.ErrorRes{
				Status: http.StatusInternalServerError,
				Titulo: "Error al Hcer la consulta",
				Cuerpo: err,
			}
			elerro.Response(w)
		} else {
			res := struct {
				Total  int
				Libros libro.ListLibros
			}{
				Total:  len(libros),
				Libros: libros,
			}
			json.NewEncoder(w).Encode(res)
		}
	}
}

func VerLibro(w http.ResponseWriter, r *http.Request) {
	Headers(&w)

	defer func() {
		if rr := recover(); rr != nil {
			w.WriteHeader(rr.(model.ErrorRes).Status)
			json.NewEncoder(w).Encode(rr)
		}
		fmt.Println("VerLibro")
	}()

	var libro libro.Libro
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])
	if err != nil {
		panic(model.ErrorRes{Titulo: "key no valida", Cuerpo: err, Mensaje: err.Error()})
	}

	libro.Ver(id)
	json.NewEncoder(w).Encode(libro)
}

func CrearLibro(w http.ResponseWriter, r *http.Request) {
	Headers(&w)

	defer func() {
		if rr := recover(); rr != nil {
			w.WriteHeader(rr.(model.ErrorRes).Status)
			json.NewEncoder(w).Encode(rr)
		}
		fmt.Println("CrearLibro")
	}()

	// Obtener datos del formulario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		// Preparar formulario y mandar la informacion
		var libro libro.LibroFormulario
		if err = json.Unmarshal(reqBody, &libro); err == nil {
			nuevoLibro, err2 := libro.Crear()
			if err != nil {
				panic(err2)
			} else {

				json.NewEncoder(w).Encode(nuevoLibro)
			}
		} else {
			panic(err)
		}

	}
}

func ActualizarLibro(w http.ResponseWriter, r *http.Request) {
	Headers(&w)

	defer func() {
		if rr := recover(); rr != nil {
			w.WriteHeader(rr.(model.ErrorRes).Status)
			json.NewEncoder(w).Encode(rr)
		}
		fmt.Println("ActualizarLibro")
	}()

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])

	if err == nil {
		reqBody, err := ioutil.ReadAll(r.Body)

		if err == nil {
			var update libro.LibroFormulario
			json.Unmarshal(reqBody, &update)

			if libro, err2 := update.Editar(id); err2 == nil {
				json.NewEncoder(w).Encode(libro)
			} else {
				panic(err2)
			}

		} else {
			json.NewEncoder(w).Encode(err.Error())
		}
	} else {
		json.NewEncoder(w).Encode(err.Error())
	}
}

func EliminarLibro(w http.ResponseWriter, r *http.Request) {
	Headers(&w)

	defer func() {
		if rr := recover(); rr != nil {
			w.WriteHeader(rr.(model.ErrorRes).Status)
			json.NewEncoder(w).Encode(rr)
		}
		fmt.Println("EliminarLibro")
	}()

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["key"])
	if err == nil {
		var l libro.Libro
		err = l.Eliminar(&id, nil)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}
	}
}
