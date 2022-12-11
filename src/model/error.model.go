package model

import (
	"encoding/json"
	"net/http"
)

type ErrorRes struct {
	Status  int    `json:"-"`
	Titulo  string `json:"titulo"`
	Mensaje string `json:"mensaje,omitempty"`
	Cuerpo  error  `json:"cuerpo,omitempty"`
}

func (t ErrorRes) Error() string {

	return t.Titulo + ": " + t.Mensaje + "\n" + t.Cuerpo.Error()
}
func (t ErrorRes) Response(w http.ResponseWriter) {
	w.WriteHeader(t.Status)
	json.NewEncoder(w).Encode(t)
}
