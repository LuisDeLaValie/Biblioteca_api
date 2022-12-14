package libro

import (
	"context"
	"fmt"
	"time"

	"libreria/src/db"
	"libreria/src/model"
	"libreria/src/model/autor"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibroPaginacions struct {
	To  int
	End int
}
type LibroOrigen struct {
	Nombre string `json:"nombre"`
	Url    string `json:"url"`
}
type Libro struct {
	Key        primitive.ObjectID `json:"key" bson:"_id,omitempty" `
	Titulo     string             `json:"titulo"`
	Sinopsis   string             `json:"Sinopsis,omitempty"`
	Autores    []*autor.Autor     `json:"autores,omitempty"`
	Editorail  string             `json:"editorial,omitempty"`
	Descargar  string             `json:"descargar,omitempty"`
	Path       string             `json:"-"`
	Verr       string             `json:"ver,omitempty"`
	Paginacion *LibroPaginacions  `json:"paginacion,omitempty"`
	Origen     *LibroOrigen       `json:"origen,omitempty"`
	Creado     *time.Time         `json:"creado,omitempty"`
}

type ListLibros []*Libro

func (l Libro) Listar(search string, all bool) (ListLibros, error) {
	var _ctx = context.Background()
	var con db.Mongodb
	defer func() {
		con.Close()
		_ctx.Done()
	}()

	libros := ListLibros{}
	var pipeline mongo.Pipeline

	matchesSearch := bson.D{}
	if search != "" {
		matchesSearch = bson.D{
			{
				Key: "$or", Value: []interface{}{
					bson.D{
						{
							Key: "titulo", Value: primitive.Regex{
								Pattern: search,
								Options: "i",
							},
						},
					},
					bson.D{
						{
							Key: "autores.nombre", Value: primitive.Regex{
								Pattern: search,
								Options: "i",
							},
						},
					},
				},
			},
		}
	}
	lookupStage := bson.D{
		{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "autor"},
				{Key: "localField", Value: "autor"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "autores"},
			},
		},
	}
	projectStage := bson.D{
		{
			Key: "$project", Value: bson.D{
				{Key: "autor", Value: false},
			},
		},
	}
	matchStage := bson.D{
		{Key: "$match", Value: matchesSearch},
	}

	if all {
		lookupStageFilter := bson.D{
			{
				Key: "$lookup", Value: bson.D{
					{Key: "from", Value: "coleccion"},
					{Key: "localField", Value: "_id"},
					{Key: "foreignField", Value: "libros"},
					{Key: "as", Value: "nada"},
				},
			},
		}
		matchStageFilter := bson.D{
			{
				Key: "$match", Value: bson.D{
					{
						Key: "nada", Value: bson.D{
							{Key: "$size", Value: 0},
						},
					},
				},
			},
		}

		pipeline = mongo.Pipeline{
			lookupStageFilter,
			matchStageFilter,
			lookupStage,
			projectStage,
			matchStage,
		}

	} else {
		pipeline = mongo.Pipeline{lookupStage, projectStage, matchStage}
	}

	if cur, err := con.GetCollection("libros").Aggregate(_ctx, pipeline); err != nil {
		return nil, err
	} else {
		for cur.Next(_ctx) {
			var aux Libro
			if err = cur.Decode(&aux); err != nil {
				return nil, err
			} else {
				libros = append(libros, &aux)
			}
		}
		return libros, nil
	}
}
func (l *Libro) Ver(key primitive.ObjectID) error {
	var _ctx = context.Background()
	var con db.Mongodb
	defer func() {
		con.Close()
		_ctx.Done()
	}()

	// filros para traer el libro
	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "autor"}, {Key: "localField", Value: "autor"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "autores"}}}}
	projectStage := bson.D{{Key: "$project", Value: bson.D{{Key: "autor", Value: false}}}}
	machtState := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: key}}}}
	limtState := bson.D{{Key: "$limit", Value: 1}}

	cur, err := con.GetCollection("libros").Aggregate(_ctx, mongo.Pipeline{lookupStage, projectStage, machtState, limtState})

	if err != nil {
		return model.ErrorRes{
			Titulo:  "Error al traer los datos",
			Mensaje: err.Error(),
			Cuerpo:  err,
		}
	}

	cur.Next(_ctx)
	err = cur.Decode(&l)

	if err != nil {
		return model.ErrorRes{
			Titulo:  "Error al obtener los datos",
			Mensaje: err.Error(),
			Cuerpo:  err,
		}

	}
	l.Descargar = "https://drive.google.com/uc?export=download&id=" + l.Path
	l.Verr = "https://drive.google.com/file/d/" + l.Path + "/view?usp=sharing"
	// https://lh3.google.com/u/0/d/1Y3dRJfPta1vlS1HjdK3I_nV7wd6yNmn1=w200-h190-p-k-nu-iv1
	l.Path = ""
	return nil
}
func (l *Libro) Eliminar(key *primitive.ObjectID, keys *[]primitive.ObjectID) error {

	var con db.Mongodb
	defer func() {
		con.Close()
	}()

	if key == nil && keys == nil {
		return fmt.Errorf("*key* y *keys* no pueden ser nil")
	} else if key != nil && keys != nil {
		return fmt.Errorf("*key* y *keys* no pueden ser diferente de nil")
	} else if keys != nil {
		filter := bson.M{
			"_id": bson.M{
				"$in": keys,
			},
		}
		_, err := con.GetCollection("libros").DeleteMany(context.TODO(), filter)
		if err != nil {
			return err
		} else {
			return nil
		}
	} else {
		filter := bson.M{"_id": key}
		_, err := con.GetCollection("libros").DeleteOne(context.TODO(), filter)
		if err != nil {
			return err
		} else {
			return nil
		}
	}

}

func (new Libro) Compare(old Libro) *string {
	mensaje := ""
	if new.Titulo != old.Titulo {
		mensaje += fmt.Sprintf("Titulo no coincide:\n\tN)%s\n\tO)%s\n", new.Titulo, old.Titulo)
	}
	if new.Sinopsis != old.Sinopsis {
		mensaje += fmt.Sprintf("Sinopsis no coincide:\n\tN)%s\n\tO)%s\n", new.Sinopsis, old.Sinopsis)
	}

	if len(new.Autores) != len(old.Autores) {
		mensaje += fmt.Sprintf("Autores no coincide:\n\tN)%d\n\tO)%d\n", len(new.Autores), len(old.Autores))
	} else {
		for i := range new.Autores {
			n := new.Autores[i]
			o := old.Autores[i]
			if *n != *o {
				mensaje += fmt.Sprintf("Autor [%d] no coincide:\n\tN)%+v\n\tO)%+v\n", i, new.Paginacion, old.Paginacion)

			}
		}
	}

	if new.Editorail != old.Editorail {
		mensaje += fmt.Sprintf("Editorail no coincide:\n\tN)%s\n\tO)%s\n", new.Editorail, old.Editorail)
	}
	if new.Descargar != old.Descargar {
		mensaje += fmt.Sprintf("Descargar no coincide:\n\tN)%s\n\tO)%s\n", new.Descargar, old.Descargar)
	}
	if new.Path != old.Path {
		mensaje += fmt.Sprintf("Path no coincide:\n\tN)%s\n\tO)%s\n", new.Path, old.Path)
	}
	if new.Verr != old.Verr {
		mensaje += fmt.Sprintf("Verr no coincide:\n\tN)%s\n\tO)%s\n", new.Verr, old.Verr)
	}

	if *new.Paginacion != *old.Paginacion {
		mensaje += fmt.Sprintf("Paginacion no coincide:\n\tN)%+v\n\tO)%+v\n", *new.Paginacion, *old.Paginacion)
	}

	if *new.Origen != *old.Origen {
		mensaje += fmt.Sprintf("Origen no coincide:\n\tN)%s\n\tO)%s\n", *new.Origen, *old.Origen)
	}

	if mensaje == "" {
		return nil
	} else {
		return &mensaje
	}
}
