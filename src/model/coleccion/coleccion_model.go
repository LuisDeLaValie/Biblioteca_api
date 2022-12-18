package coleccion

import (
	"context"
	"fmt"
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
	Sinopsis   string             `json:"sinopsis,omitempty" bson:"sinopsis,omitempty"`
	Libos_list []libro.Libro      `json:"libros,omitempty" bson:"libros"`
	// Path       string             `json:"path,omitempty" bson:"path,omitempty"`
	Creado time.Time `json:"creado" bson:"creado"`
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
	lookupStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "libros"},
			{Key: "localField", Value: "libros"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "libros"},
		}},
	}
	projectStage := bson.D{
		{
			Key: "$project", Value: bson.D{
				{Key: "libros.titulo", Value: true},
				{Key: "libros._id", Value: true},
			},
		},
	}
	matchStage := bson.D{
		{Key: "$match", Value: matchesSearch},
	}
	cur, err := con.GetCollection("coleccion").Aggregate(_ctx, mongo.Pipeline{lookupStage, projectStage, matchStage})

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
				{Key: "localField", Value: "libros"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "libros"},
			},
		},
	}
	machtState := bson.D{
		{
			Key: "$match", Value: bson.D{
				{Key: "_id", Value: key},
			},
		},
	}
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

func (coll Coleccion) Eliminar(key primitive.ObjectID, all bool) error {

	var con db.Mongodb
	defer func() {
		con.Close()
	}()

	filter := bson.M{"_id": key}
	if all {
		var form libro.LibroFormulario
		// db.coleccion.find({_id:ObjectId('6371726c7a56105ea05dab3f')},{libros:1,_id:0})
		val := con.GetCollection("coleccion").FindOne(context.TODO(), filter)
		if err := val.Decode(&form); err != nil {
			return err
		} else {
			var li libro.Libro
			var aux ColeccionFormulario
			val.Decode(&aux)
			if err := li.Eliminar(nil, &aux.Libors); err != nil {
				return err
			}
		}
	}
	_, err := con.GetCollection("coleccion").DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	} else {
		return nil
	}
}
func (l Coleccion) Compare(com Coleccion) *string {
	mensaje := ""

	if l.Titulo != com.Titulo {
		mensaje += fmt.Sprintf("Titulo no coincide:\n\tN)%s\n\tO)%s\n", l.Titulo, com.Titulo)
	}
	if l.Sinopsis != com.Sinopsis {
		mensaje += fmt.Sprintf("Sinopsis no coincide:\n\tN)%s\n\tO)%s\n", l.Sinopsis, com.Sinopsis)
	}
	if len(l.Libos_list) != len(com.Libos_list) {
		mensaje += fmt.Sprintf("Libos_list no coincide:\n\tN)%d\n\tO)%d\n", len(l.Libos_list), len(com.Libos_list))
	} else {
		for i := range l.Libos_list {
			n := l.Libos_list[i]
			o := com.Libos_list[i]
			if val := n.Compare(o); val != nil {
				mensaje += fmt.Sprintf("libros no coincide:\n\tN)[%d]%+v\n\tO)[%d]%+v\n", i, n, i, o)

			}
		}
	}

	if mensaje == "" {
		return nil
	} else {
		return &mensaje
	}
}
