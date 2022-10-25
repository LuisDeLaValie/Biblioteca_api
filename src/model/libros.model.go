package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	// "fmt"

	conn "libreria/src/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type _Paginacions struct {
	To  int
	End int
}

type Libro struct {
	Key        primitive.ObjectID   `bson:"_id,omitempty" json:"key,omitempty"`
	Titulo     string               `json:"titulo"`
	Sinopsis   string               `json:"sipnosis"`
	Autor_DB   []primitive.ObjectID `bson:"autor,omitempty" json:"ref_autor,omitempty"`
	Autor_list []Autor              `bson:"A_l,omitempty" json:"autor,omitempty"`
	Editorail  string               `json:"editorial"`
	Paginacion _Paginacions         `json:"paginacion"`
	Origen     origen               `json:"origen"`
	Path       string               `json:"path"`
	Creado     time.Time            `json:"creado"`
}

type origen struct {
	Nombre string `json:"nombre"`
	Url    string `json:"url"`
}

type ListLibros []*Libro

var ctx = context.Background()
var collecion = conn.GetCollection("libros")

// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/

func (this *Libro) Crear() error {

	this.Creado = time.Now()
	_, err := collecion.InsertOne(ctx, this)

	if err != nil {
		return err
	} else {
		return nil
	}
}
func (this *Libro) Listar() (ListLibros, error) {

	var libros ListLibros
	filter := bson.D{}
	// *mongo.InsertOneResult

	cur, err := collecion.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var aux Libro
		err = cur.Decode(&aux)

		if err != nil {
			return nil, err
		}

		libros = append(libros, &aux)
	}

	return libros, nil

}
func (this *Libro) Ver() error {

	filter := bson.D{{"_id", this.Key}}

	err := collecion.FindOne(ctx, filter).Decode(&this)
	if err != nil {
		return err
	} else {
		return nil
	}

}
func (this *Libro) Editar() error {
	// oid,_ := primitive.ObjectIDFromHex(this.Key)

	filtter := bson.M{"_id": this.Key}
	update := bson.M{
		"$set": bson.M{
			"titulo": this.Titulo,
			"path":   this.Path,
		},
	}

	_, err := collecion.UpdateOne(ctx, filtter, update)

	if err != nil {
		return err
	}

	return nil
}
func (this *Libro) Eliminar() error {
	filter := bson.D{{"_id", this.Key}}
	_, err := collecion.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
