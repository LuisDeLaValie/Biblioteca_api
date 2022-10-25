package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)



type Coleccion struct {
	Key primitive.ObjectID `bson:"_id,omitempty" json:"key,omitempty"`
	Titulo string `json:"titulo"`
	Sinopsis string `json:"sipnosis"`
	libors_db []primitive.ObjectID `bson:"libros,omitempty" json:"libros_db,omitempty"`
	Libos_list []Libro `bosn:"L_s,omitempty" json:"libros,omitempty"`
	Path string `json:"path"`
	Creado time.Time `json:"creado"`
}

func (this *Coleccion) Crear() error {
	return nil
}
func (this *Coleccion) Listar() error {
	return nil
}
func (this *Coleccion) Ver() error {
	return nil
}
func (this *Coleccion) Editar() error {
	return nil
}
func (this *Coleccion) Actualizar() error {
	return nil
}
func (this *Coleccion) eliminar() error {
	return nil
}