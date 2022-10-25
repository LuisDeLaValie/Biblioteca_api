package handler
import 	(
	"net/http"
	"encoding/json"
	"io/ioutil"
	
	m "libreria/src/model"

	// "github.com/gorilla/mux"

)

type ErrorRes struct{
	Error string `json:"error"`
	Mensaje string `json:"mensaje,omitempty"`
	Cuerpo error `json:"cuerpo,omitempty"`
}


func ListarLibros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var aux  m.Libro
	libros, err:= aux.Listar()

	if err != nil {				
		cerror:=ErrorRes{Error:"Error obteniendo los datos", Cuerpo:err,Mensaje:err.Error()}
		json.NewEncoder(w).Encode(cerror)
	}
	json.NewEncoder(w).Encode(libros)

	
}
func VerLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var newLibro m.Libro
	reqBody,err := ioutil.ReadAll(r.Body)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	json.Unmarshal(reqBody,&newLibro)
	err = newLibro.Ver()

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return;
	}

	json.NewEncoder(w).Encode(newLibro)
}
func CrearLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var newLibro m.Libro
	reqBody,err := ioutil.ReadAll(r.Body)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	json.Unmarshal(reqBody,&newLibro)
	err = newLibro.Crear()

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	json.NewEncoder(w).Encode(newLibro)
}
func ActualizarLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var newLibro m.Libro
	reqBody,err := ioutil.ReadAll(r.Body)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	json.Unmarshal(reqBody,&newLibro)
	err = newLibro.Editar()

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	json.NewEncoder(w).Encode("Libro actualizado con exito")
}
func EliminarLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var newLibro m.Libro
	reqBody,err := ioutil.ReadAll(r.Body)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	json.Unmarshal(reqBody,&newLibro)
	err=newLibro.Eliminar()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode("delete")
}