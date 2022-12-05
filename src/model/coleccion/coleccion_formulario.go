package coleccion

import (
	"context"
	"libreria/src/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ColeccionFormulario struct {
	Titulo   string               `json:"titulo" bson:"titulo"`
	Sinopsis string               `json:"sinopsis,omitempty" bson:"sinopsis,omitempty"`
	Libors   []primitive.ObjectID `json:"libros" bson:"libros"`
	Path     string               `json:"path,omitempty" bson:"path,omitempty"`
	Creado   time.Time            `json:"creado,omitempty" bson:"creado"`
}

func (coleccion *ColeccionFormulario) Crear() (*Coleccion, error) {
	var _ctx = context.Background()
	var con db.Mongodb
	defer func() {
		con.Close()
		_ctx.Done()
	}()
	// Preparar datos del formulario
	coleccion.Creado = time.Now()

	// Insertar Libro
	oid, err := con.GetCollection("coleccion").InsertOne(_ctx, coleccion)

	if err == nil {

		id := oid.InsertedID.(primitive.ObjectID)
		var col Coleccion
		err := col.Ver(id)

		if err != nil {
			return nil, err
		} else {
			return &col, nil
		}

	} else {
		return nil, err
	}
}

func (upColecc *ColeccionFormulario) Editar(key primitive.ObjectID) (*Coleccion, error) {
	var _ctx = context.Background()

	var con db.Mongodb
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

	_, err := con.GetCollection("coleccion").UpdateOne(_ctx, filtter, update)
	if err != nil {
		return nil, err
	} else {
		var coleccion Coleccion
		err := coleccion.Ver(key)

		if err != nil {
			return nil, err
		} else {
			return &coleccion, nil
		}
	}

}
