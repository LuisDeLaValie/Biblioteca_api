package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	conn "libreria/src/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/

type _Paginacions struct {
	To  int
	End int
}
type origen struct {
	Nombre string `json:"nombre"`
	Url    string `json:"url"`
}
type Libro struct {
	Key        primitive.ObjectID `json:"key" bson:"_id,omitempty" `
	Titulo     string             `json:"titulo"`
	Sinopsis   string             `json:"sipnosis,omitempty"`
	Autores    []Autor            `json:"autores,omitempty"`
	Editorail  string             `json:"editorial,omitempty"`
	Paginacion _Paginacions       `json:"paginacion,omitempty"`
	Origen     origen             `json:"origen,omitempty"`
	Path       string             `json:"path,omitempty"`
	Creado     time.Time          `json:"creado"`
}

type ListLibros []*Libro

/// Listar todos los sibros
func (l Libro) Listar() (ListLibros, error) {
	var _ctx = context.Background()
	var _collecion = conn.GetCollection("libros")

	var libros ListLibros

	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "autor"}, {Key: "localField", Value: "autor"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "autores"}}}}
	projectStage := bson.D{{Key: "$project", Value: bson.D{{Key: "autor", Value: false}}}}

	cur, err := _collecion.Aggregate(_ctx, mongo.Pipeline{lookupStage, projectStage})

	if err == nil {
		for cur.Next(_ctx) {
			var aux Libro
			err = cur.Decode(&aux)

			if err != nil {
				return nil, err
			}

			libros = append(libros, &aux)
		}

		return libros, nil
	} else {
		return nil, err
	}
}
func (l *Libro) Ver(key primitive.ObjectID) error {
	var _ctx = context.Background()
	var _collecion = conn.GetCollection("libros")

	// filros para traer el libro
	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "autor"}, {Key: "localField", Value: "autor"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "autores"}}}}
	projectStage := bson.D{{Key: "$project", Value: bson.D{{Key: "autor", Value: false}}}}
	machtState := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: key}}}}
	limtState := bson.D{{Key: "$limit", Value: 1}}

	cur, err := _collecion.Aggregate(_ctx, mongo.Pipeline{lookupStage, projectStage, machtState, limtState})

	if err == nil {
		cur.Next(_ctx)
		err = cur.Decode(&l)
		if err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return err
	}
}
func (l *Libro) Eliminar(key primitive.ObjectID) error {

	var _collecion = conn.GetCollection("libros")

	filter := bson.M{"_id": key}
	_, err := _collecion.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	} else {
		return nil
	}
}

type LibroFormulario struct {
	Titulo     string               `json:"titulo" bson:"titulo"`
	Sipnosis   string               `json:"sinopsis,omitempty" bson:"sinopsis,omitempty"`
	Autores    []primitive.ObjectID `json:"autores,omitempty" bson:"autor,omitempty"`
	Editorail  string               `json:"editorial,omitempty" bson:"editorial,omitempty"`
	Paginas    int                  `json:"paginas,omitempty" bson:"-"`
	Pagina     int                  `json:"pagina,omitempty" bson:"-"`
	Origen     origen               `json:"origen" bson:"origen"`
	Path       string               `json:"path,omitempty" bson:"path,omitempty"`
	Creado     time.Time            `json:"creado,omitempty" bson:"creado,omitempty"`
	Paginacion _Paginacions         `json:"_,omitempty" bson:"paginacion,omitempty"`
}

func (libro *LibroFormulario) Crear() (*Libro, error) {
	var _ctx = context.Background()
	var _collecion = conn.GetCollection("libros")

	// Preparar datos del formulario
	libro.Creado = time.Now()
	libro.Paginacion = _Paginacions{0, libro.Paginas}

	// Insertar Libro
	oid, err := _collecion.InsertOne(_ctx, libro)

	if err == nil {

		var nuevoLibro *Libro
		err := nuevoLibro.Ver(oid.InsertedID.(primitive.ObjectID))

		if err != nil {
			return nil, err
		} else {
			return nuevoLibro, nil
		}

	} else {
		return nil, err
	}
}
func (upLibro *LibroFormulario) Editar(key primitive.ObjectID) (*Libro, error) {
	var _ctx = context.Background()
	var _collecion = conn.GetCollection("libros")

	filtter := bson.M{"_id": key}
	update := bson.M{
		"$set": bson.M{
			"titulo":        upLibro.Titulo,
			"sinopsis":      upLibro.Sipnosis,
			"autor":         upLibro.Autores,
			"editorial":     upLibro.Editorail,
			"origen":        upLibro.Origen,
			"path":          upLibro.Path,
			"paginacion.to": upLibro.Pagina,
		},
	}

	_, err := _collecion.UpdateOne(_ctx, filtter, update)
	if err != nil {
		return nil, err
	} else {

		var libro *Libro
		err := libro.Ver(key)

		if err != nil {
			return nil, err
		} else {
			return libro, nil
		}
	}
}
