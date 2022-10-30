package model

import (
	"context"
	conn "libreria/src/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Autor struct {
	Key    primitive.ObjectID `bson:"_id,omitempty" json:"key,omitempty"`
	Nombre string             `bson:"nombre" json:"nombre"`
}

type ListAutor []*Autor

func CrearAutor(autor Autor) (*Autor, error) {
	var _ctx = context.Background()
	var _collecion = conn.GetCollection("autor")

	// Insertar Libro
	oid, err := _collecion.InsertOne(_ctx, autor)

	if err == nil {
		nuevoAutor := VerAutor(oid.InsertedID.(primitive.ObjectID))
		return nuevoAutor, nil

	} else {
		return nil, err
	}
}

/// Listar todos los sibros
func ListarAutor() (ListAutor, error) {
	var _ctx = context.Background()
	var _collecion = conn.GetCollection("autor")

	var autores ListAutor

	filter := bson.M{}
	cur, err := _collecion.Find(_ctx, filter)

	if err == nil {
		for cur.Next(_ctx) {
			var aux Autor
			err = cur.Decode(&aux)

			if err != nil {
				return nil, err
			}
			autores = append(autores, &aux)
		}
		return autores, nil
	} else {
		return nil, err
	}
}

func VerAutor(key primitive.ObjectID) *Autor {
	var _ctx = context.Background()
	var _collecion = conn.GetCollection("autor")

	filter := bson.M{"_id": key}
	result := _collecion.FindOne(_ctx, filter)

	var autor *Autor
	result.Decode(&autor)
	return autor
}

func EditarAutor(key primitive.ObjectID, upAutor Autor) (*Autor, error) {
	var _ctx = context.Background()
	var _collecion = conn.GetCollection("autor")

	filtter := bson.M{"_id": key}
	update := bson.M{
		"$set": bson.M{
			"nombre": upAutor.Nombre,
		},
	}

	_, err := _collecion.UpdateOne(_ctx, filtter, update)
	if err != nil {
		return nil, err
	} else {
		autor := VerAutor(key)
		return autor, nil
	}

}
func EliminarAutor(key primitive.ObjectID) error {

	var _collecion = conn.GetCollection("autor")

	filter := bson.M{"_id": key}
	_, err := _collecion.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	} else {
		return nil
	}
}
