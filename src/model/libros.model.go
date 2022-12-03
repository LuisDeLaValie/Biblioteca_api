package model

import (
	"context"
	"fmt"
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
	Descargar  string             `json:"descargar,omitempty"`
	Path       string             `json:"-"`
	Verr       string             `json:"ver,omitempty"`
	Paginacion _Paginacions       `json:"paginacion,omitempty"`
	Origen     origen             `json:"origen,omitempty"`
	Creado     time.Time          `json:"creado"`
}

type ListLibros []*Libro

/// Listar todos los sibros
func (l Libro) Listar(search string, all bool) ListLibros {
	var _ctx = context.Background()
	var con conn.Mongodb
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

	if !all {
		pipeline = mongo.Pipeline{lookupStage, projectStage, matchStage}
	} else {
		// left joint
		lookupStageFilter := bson.D{
			{
				Key: "$lookup", Value: bson.D{
					{Key: "from", Value: "coleccion"},
					{Key: "localField", Value: "_id"},
					{Key: "foreignField", Value: "libros.id"},
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
		pipeline = mongo.Pipeline{lookupStageFilter, matchStageFilter, lookupStage, projectStage, matchStage}
	}

	cur, err := con.GetCollection("libros").Aggregate(_ctx, pipeline)

	if err != nil {
		panic(ErrorRes{Cuerpo: err, Mensaje: err.Error()})
	}
	for cur.Next(_ctx) {

		var aux Libro
		err = cur.Decode(&aux)

		if err != nil {
			panic(ErrorRes{Cuerpo: err, Mensaje: err.Error()})
		}
		libros = append(libros, &aux)
	}

	return libros
}
func (l *Libro) Ver(key primitive.ObjectID) {
	var _ctx = context.Background()
	var con conn.Mongodb
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
		panic(ErrorRes{Error: "Error al buscar datos", Cuerpo: err, Mensaje: err.Error()})
	}

	cur.Next(_ctx)
	err = cur.Decode(&l)

	if err != nil {
		panic(ErrorRes{Error: "Error al obtener los datos", Cuerpo: err, Mensaje: err.Error()})
	}
	l.Descargar = "https://drive.google.com/uc?export=download&id=" + l.Path
	l.Verr = "https://drive.google.com/file/d/" + l.Path + "/view?usp=sharing"
	// https://lh3.google.com/u/0/d/1Y3dRJfPta1vlS1HjdK3I_nV7wd6yNmn1=w200-h190-p-k-nu-iv1
}
func (l *Libro) Eliminar(key primitive.ObjectID) error {

	var con conn.Mongodb
	defer func() {
		con.Close()
	}()
	filter := bson.M{"_id": key}
	_, err := con.GetCollection("libros").DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// metodo de descarga
/* func (l *Libro) Descargar(key primitive.ObjectID) []byte {
	l.Ver(key)
	///Volumes/GoogleDrive/Mi unidad/Libros/overlord/Overlord Volumen 9.pdf
	// path := "/Volumes/GoogleDrive/Mi unidad/Libros" + l.Path
	path := "/Volumes/GoogleDrive/Mi unidad/Libros/overlord/Overlord Volumen 9.pdf"

	/// tambien teneenmos el paquete bufio para leer archivos
	archivo, err := os.Open(path)
	// difer nos ayuda a que sin importa lo que pase se serrara el
	// archi para no mantenerlo abierto
	defer archivo.Close()

	if err != nil {
		panic(ErrorRes{Cuerpo: err, Mensaje: err.Error()})
	}
	scanner := bufio.NewScanner(archivo)
	return scanner.Bytes()

	// if fileBytes, err := ioutil.ReadFile(path); err == nil {
	// 	return fileBytes
	// } else {

	// 	panic(ErrorRes{Cuerpo: err, Mensaje: err.Error()})
	// }

} */

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
	Paginacion _Paginacions         `json:"-" bson:"paginacion,omitempty"`
}

func (libro *LibroFormulario) Crear() Libro {
	var con conn.Mongodb
	defer func() {
		con.Close()
	}()
	var nuevoLibro Libro

	// Preparar datos del formulario
	libro.Creado = time.Now()
	libro.Paginacion = _Paginacions{0, libro.Paginas}

	fmt.Print(libro)

	// Insertar Libro
	oid, err := con.GetCollection("libros").InsertOne(context.TODO(), libro)
	if err != nil {
		panic(ErrorRes{Cuerpo: err, Mensaje: err.Error()})
	}

	nuevoLibro.Ver(oid.InsertedID.(primitive.ObjectID))
	return nuevoLibro
}

func (upLibro *LibroFormulario) Editar(key primitive.ObjectID) Libro {
	var con conn.Mongodb
	defer func() {
		con.Close()
	}()
	var libro Libro

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

	_, err := con.GetCollection("libros").UpdateOne(context.TODO(), filtter, update)
	if err != nil {
		panic(ErrorRes{Cuerpo: err, Mensaje: err.Error()})
	}

	libro.Ver(key)
	return libro
}
