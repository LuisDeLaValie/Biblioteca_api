package libro

import (
	"context"
	"fmt"
	"time"

	"libreria/src/db"
	"libreria/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LibroFormulario struct {
	Titulo     string               `json:"titulo" bson:"titulo"`
	Sipnosis   string               `json:"sinopsis,omitempty" bson:"sinopsis,omitempty"`
	Autores    []primitive.ObjectID `json:"autores,omitempty" bson:"autor,omitempty"`
	Editorail  string               `json:"editorial,omitempty" bson:"editorial,omitempty"`
	Paginas    int                  `json:"paginas,omitempty" bson:"-"`
	Pagina     int                  `json:"pagina,omitempty" bson:"-"`
	Origen     LibroOrigen          `json:"origen" bson:"origen"`
	Path       string               `json:"path,omitempty" bson:"path,omitempty"`
	Creado     time.Time            `json:"creado,omitempty" bson:"creado,omitempty"`
	Paginacion LibroPaginacions     `json:"-" bson:"paginacion,omitempty"`
}

func (libro *LibroFormulario) Crear() (*Libro, error) {
	var con db.Mongodb
	defer func() {
		con.Close()
	}()
	var nuevoLibro Libro

	// Preparar datos del formulario
	libro.Creado = time.Now()
	libro.Paginacion = LibroPaginacions{0, libro.Paginas}

	// Insertar Libro
	oid, err := con.GetCollection("libros").InsertOne(context.TODO(), libro)
	if err != nil {
		return nil, err
	}

	nuevoLibro.Ver(oid.InsertedID.(primitive.ObjectID))
	return &nuevoLibro, nil
}

func (upLibro *LibroFormulario) Editar(key primitive.ObjectID) (*Libro, error) {
	var con db.Mongodb
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
		return nil, model.ErrorRes{Cuerpo: err, Mensaje: err.Error()}

	}

	libro.Ver(key)
	return &libro, nil
}

func (l LibroFormulario) Compare(com LibroFormulario) *string {
	mensaje := ""
	if l.Titulo != com.Titulo {
		mensaje += fmt.Sprintf("Titulo no coincide:\n\t%s\n\t%s", l.Titulo, com.Titulo)
	}
	if l.Sipnosis != com.Sipnosis {
		mensaje += fmt.Sprintf("Sipnosis no coincide:\n\t%s\n\t%s", l.Sipnosis, com.Sipnosis)
	}

	if len(l.Autores) != len(com.Autores) {
		mensaje += fmt.Sprintf("Autores no coincide:\n\t%d\n\t%d", len(l.Autores), len(com.Autores))
	} else {
		for i := range l.Autores {
			n := l.Autores[i]
			o := com.Autores[i]
			if n != o {
				mensaje += fmt.Sprintf("Autores no coincide:\n\t[%d]%s\n\t[%d]%s", i, n, i, o)
			}
		}
	}
	if l.Editorail != com.Editorail {
		mensaje += fmt.Sprintf("Editorail no coincide:\n\t%s\n\t%s", l.Editorail, com.Editorail)
	}
	if l.Paginas != com.Paginas {
		mensaje += fmt.Sprintf("Paginas no coincide:\n\t%d\n\t%d", l.Paginas, com.Paginas)
	}
	if l.Pagina != com.Pagina {
		mensaje += fmt.Sprintf("Pagina no coincide:\n\t%d\n\t%d", l.Pagina, com.Pagina)
	}
	if l.Origen != com.Origen {
		mensaje += fmt.Sprintf("Origen no coincide:\n\t%s\n\t%s", l.Origen, com.Origen)
	}
	if l.Path != com.Path {
		mensaje += fmt.Sprintf("Path no coincide:\n\t%s\n\t%s", l.Path, com.Path)
	}
	if l.Creado != com.Creado {
		mensaje += fmt.Sprintf("Creado no coincide:\n\t%s\n\t%s", l.Creado, com.Creado)
	}

	if mensaje == "" {
		return nil
	} else {
		return &mensaje
	}
}
