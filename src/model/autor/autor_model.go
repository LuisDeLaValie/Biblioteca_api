package autor

import (
	"context"
	"libreria/src/db"

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
	var con db.Mongodb
	defer func() {
		con.Close()
		_ctx.Done()
	}()
	// Insertar Libro
	oid, err := con.GetCollection("autor").InsertOne(_ctx, autor)

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
	var con db.Mongodb
	defer func() {
		con.Close()
		_ctx.Done()
	}()
	var autores ListAutor

	matchesSearch := bson.D{}
	if search != "" {
		matchesSearch = bson.D{
			{
				Key: "nombre", Value: primitive.Regex{
					Pattern: search,
					Options: "i",
				},
			},
		}
	}

	cur, err := con.GetCollection("autor").Find(_ctx, matchesSearch)

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
func (autor *Autor) Ver(key primitive.ObjectID) {
	var _ctx = context.Background()
	var con db.Mongodb
	defer func() {
		con.Close()
		_ctx.Done()
	}()
	filter := bson.M{"_id": key}
	result := con.GetCollection("autor").FindOne(_ctx, filter)

	result.Decode(&autor)
}
func (upAutor *Autor) Editar(key primitive.ObjectID) error {
	var _ctx = context.Background()
	var con db.Mongodb
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

	_, err := con.GetCollection("autor").UpdateOne(_ctx, filtter, update)
	if err != nil {
		return err
	} else {
		upAutor.Key = key
		return nil
	}
}
func (autor Autor) Eliminar(key primitive.ObjectID) error {

	var con db.Mongodb
	defer func() {
		con.Close()
	}()

	filter := bson.M{"_id": key}
	_, err := con.GetCollection("autor").DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	} else {
		return nil
	}
}
