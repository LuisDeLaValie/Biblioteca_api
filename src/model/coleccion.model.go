package model

import (
	"context"
	"time"

	conn "libreria/src/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Coleccion struct {
	Key        primitive.ObjectID `json:"key,omitempty" bson:"_id,omitempty"`
	Titulo     string             `json:"titulo" bson:"titulo"`
	Sinopsis   string             `json:"sipnosis,omitempty" bson:"sipnosis,omitempty"`
	Libos_list []Libro            `json:"libros,omitempty" bson:"libros"`
	Path       string             `json:"path,omitempty" bson:"path,omitempty"`
	Creado     time.Time          `json:"creado" bson:"creado"`
}
type ListColeccion []*Coleccion

/// Listar todos los sibros
func (coll Coleccion) Listar(search string) (ListColeccion, error) {
	var _ctx = context.Background()
	var con conn.Mongodb
	var _collecion = con.GetCollection("coleccion")
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
	cur, err := _collecion.Aggregate(_ctx, mongo.Pipeline{lookupStage, matchStage})

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
	var con conn.Mongodb
	var _collecion = con.GetCollection("coleccion")
	defer func() {
		con.Close()
		_ctx.Done()
	}()
	// filros para traer el libro
	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "libros"}, {Key: "localField", Value: "libros"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "libros"}}}}
	machtState := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: key}}}}
	limtState := bson.D{{Key: "$limit", Value: 1}}

	cur, err := _collecion.Aggregate(_ctx, mongo.Pipeline{lookupStage, machtState, limtState})

	if err == nil {
		var coleccion *Coleccion
		cur.Next(_ctx)
		err = cur.Decode(&coleccion)

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

	var con conn.Mongodb
	var _collecion = con.GetCollection("coleccion")
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

type ColeccionFormulario struct {
	Titulo   string               `json:"titulo" bson:"titulo"`
	Sinopsis string               `json:"sinopsis,omitempty" bson:"sinopsis,omitempty"`
	Libors   []primitive.ObjectID `json:"libros" bson:"libros"`
	Path     string               `json:"path,omitempty" bson:"path,omitempty"`
	Creado   time.Time            `json:"creado,omitempty" bson:"creado"`
}

func (coleccion *ColeccionFormulario) Crear() (*Coleccion, error) {
	var _ctx = context.Background()
	var con conn.Mongodb
	var _collecion = con.GetCollection("coleccion")
	defer func() {
		con.Close()
		_ctx.Done()
	}()
	// Preparar datos del formulario
	coleccion.Creado = time.Now()

	// Insertar Libro
	oid, err := _collecion.InsertOne(_ctx, coleccion)

	if err == nil {
		var nuevaColeccion *Coleccion
		err := nuevaColeccion.Ver(oid.InsertedID.(primitive.ObjectID))

		if err != nil {
			return nil, err
		} else {
			return nuevaColeccion, nil
		}

	} else {
		return nil, err
	}
}

func (upColecc *ColeccionFormulario) Editar(key primitive.ObjectID) (*Coleccion, error) {
	var _ctx = context.Background()

	var con conn.Mongodb
	var _collecion = con.GetCollection("coleccion")
	defer func() {
		con.Close()
		_ctx.Done()
	}()
	filtter := bson.M{"_id": key}
	update := bson.M{
		"$set": bson.M{
			"titulo":   upColecc.Titulo,
			"sinopsis": upColecc.Sinopsis,
			"libros":   upColecc.Libors,
			"path":     upColecc.Path,
		},
	}

	_, err := _collecion.UpdateOne(_ctx, filtter, update)
	if err != nil {
		return nil, err
	} else {
		var coleccion *Coleccion
		err := coleccion.Ver(key)

		if err != nil {
			return nil, err
		} else {
			return coleccion, nil
		}
	}

}
