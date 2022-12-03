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

func (autor *Autor) Crear() error {
	var _ctx = context.Background()
	var con conn.Mongodb
	var _collecion = con.GetCollection("autor")
	defer func() {
		con.Close()
		_ctx.Done()
	}()
	// Insertar Libro
	oid, err := _collecion.InsertOne(_ctx, autor)

	if err == nil {
		autor.Key = oid.InsertedID.(primitive.ObjectID)
		return nil
	} else {
		return err
	}
}

/// Listar todos los sibros
func (autor Autor) Listar(search string) (ListAutor, error) {
	var _ctx = context.Background()
	var con conn.Mongodb
	var _collecion = con.GetCollection("autor")
	defer func() {
		con.Close()
		_ctx.Done()
	}()
	var autores ListAutor

	matchesSearch := bson.D{}
	if search != "" {
		matchesSearch = bson.D{
			{Key: "titulo", Value: primitive.Regex{
				Pattern: search,
				Options: "i",
			}},
		}
	}
	matchStage := bson.D{
		{Key: "$match", Value: matchesSearch},
	}
	cur, err := _collecion.Find(_ctx, matchStage)

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
func (autor Autor) Ver(key primitive.ObjectID) {
	var _ctx = context.Background()
	var con conn.Mongodb
	var _collecion = con.GetCollection("autor")
	defer func() {
		con.Close()
		_ctx.Done()
	}()
	filter := bson.M{"_id": key}
	result := _collecion.FindOne(_ctx, filter)

	result.Decode(&autor)
}
func (upAutor *Autor) Editar(key primitive.ObjectID) error {
	var _ctx = context.Background()
	var con conn.Mongodb
	var _collecion = con.GetCollection("autor")
	defer func() {
		con.Close()
		_ctx.Done()
	}()
	filtter := bson.M{"_id": key}
	update := bson.M{
		"$set": bson.M{
			"nombre": upAutor.Nombre,
		},
	}

	_, err := _collecion.UpdateOne(_ctx, filtter, update)
	if err != nil {
		return err
	} else {
		upAutor.Key = key
		return nil
	}
}
func (autor Autor) Eliminar(key primitive.ObjectID) error {

	var con conn.Mongodb
	var _collecion = con.GetCollection("autor")
	defer func() {
		con.Close()
	}()

	filter := bson.M{"_id": key}
	_, err := _collecion.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	} else {
		return nil
	}
}
