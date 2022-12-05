package coleccion

import (
	"context"
	"time"

	"libreria/src/db"
	"libreria/src/model/libro"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Coleccion struct {
	Key        primitive.ObjectID `json:"key,omitempty" bson:"_id,omitempty"`
	Titulo     string             `json:"titulo" bson:"titulo"`
	Sinopsis   string             `json:"sipnosis,omitempty" bson:"sipnosis,omitempty"`
	Libos_list []libro.Libro      `json:"libros,omitempty" bson:"libros"`
	Path       string             `json:"path,omitempty" bson:"path,omitempty"`
	Creado     time.Time          `json:"creado" bson:"creado"`
}
type ListColeccion []*Coleccion

/// Listar todos los sibros
func (coll Coleccion) Listar(search string) (ListColeccion, error) {
	var _ctx = context.Background()
	var con db.Mongodb
	defer func() {
		con.Close()
		_ctx.Done()
	}()
	var colecciones ListColeccion
	matchesSearch := bson.D{}
	if search != "" {
		matchesSearch = bson.D{
			{Key: "titulo", Value: primitive.Regex{
				Pattern: search,
				Options: "i",
			}},
		}
	}
	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "libros"}, {Key: "localField", Value: "libros"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "libros"}}}}
	matchStage := bson.D{
		{Key: "$match", Value: matchesSearch},
	}
	cur, err := con.GetCollection("coleccion").Aggregate(_ctx, mongo.Pipeline{lookupStage, matchStage})

	if err == nil {
		for cur.Next(_ctx) {
			var aux Coleccion
			err = cur.Decode(&aux)

			if err != nil {
				return nil, err
			}

			colecciones = append(colecciones, &aux)
		}

		return colecciones, nil
	} else {
		return nil, err
	}
}

func (coll *Coleccion) Ver(key primitive.ObjectID) error {
	var _ctx = context.Background()
	var con db.Mongodb
	defer func() {
		con.Close()
		_ctx.Done()
	}()
	// filros para traer el libro
	lookupStage := bson.D{
		{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "libros"},
				{Key: "localField", Value: "libros.id"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "libros"},
			},
		},
	}
	machtState := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: key}}}}
	limtState := bson.D{{Key: "$limit", Value: 1}}

	cur, err := con.GetCollection("coleccion").Aggregate(_ctx, mongo.Pipeline{lookupStage, machtState, limtState})

	if err == nil {
		cur.Next(_ctx)
		err = cur.Decode(&coll)

		if err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func (coll Coleccion) Eliminar(key primitive.ObjectID) error {

	var con db.Mongodb
	defer func() {
		con.Close()
	}()
	filter := bson.M{"_id": key}
	_, err := con.GetCollection("coleccion").DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	} else {
		return nil
	}
}
