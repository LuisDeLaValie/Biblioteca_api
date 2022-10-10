package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)


type CollectionLibros struct {
	Key primitive.ObjectID `bson:"_id,omitempty" json:"key,omitempty"`
	Titulo string `json:"titulo"`
	ListaLibros []primitive.ObjectID `bson:"libros,omitempty" json:"keyLibros,omitempty"`
	Libros []Libros `bson:"_id,omitempty" json:"libros"`
	Creado time.Time `json:"creado"`
}

func (this *CollectionLibros) Crear() error {
	return nil
}
func (this *CollectionLibros) Listar() error {
	return nil
}
func (this *CollectionLibros) Ver() error {
	return nil
}
func (this *CollectionLibros) Editar() error {
	return nil
}
func (this *CollectionLibros) Actualizar() error {
	return nil
}
func (this *CollectionLibros) eliminar() error {
	return nil
}