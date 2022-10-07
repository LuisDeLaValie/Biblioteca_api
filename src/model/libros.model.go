package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"context"
	// "fmt"

 	conn "libreria/src/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


 


type Libro struct {
	Key primitive.ObjectID `bson:"_id,omitempty" json:"key,omitempty"`
	Titulo string `json:"titulo"`
	Path string `json:"path"`
	Creado time.Time `json:"creado"`

}

type ListLibros []*Libro


var 	ctx = 	context.Background()
var collecion = conn.GetCollection("lirbos")

func (this *Libro) Crear()  error {
		
	this.Creado = time.Now()
	_, err := collecion.InsertOne(ctx,this)

	if err != nil {
		return  err
	}else {
		return nil
	}
}
func (this *Libro) Listar() (ListLibros, error) {
		
	var libros ListLibros
	filter := bson.D{}
	// *mongo.InsertOneResult

	cur, err := collecion.Find(ctx,filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var aux Libro
		err = cur.Decode(&aux)

		if err != nil {
			return nil, err
		}

		libros = append(libros,&aux)
	}

	return libros, nil

}
func (this *Libro) Ver()  error {
	return nil
}
func (this *Libro) Editar() error {
	// oid,_ := primitive.ObjectIDFromHex(this.Key)
	
	filtter := bson.M{"_id": this.Key}
	update := bson.M{
		"$set":bson.M{
			"titulo":this.Titulo,
			"path":this.Path,
		},
	}

	_,err := collecion.UpdateOne(ctx,filtter,update)
	
	if err != nil {
		return err
	}

	return nil
}
func (this *Libro) Actualizar() error {
	return nil
}
func (this *Libro) eliminar() error {
	return nil
}
