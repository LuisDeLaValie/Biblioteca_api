package autor_test

import (
	"libreria/src/model/autor"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var id primitive.ObjectID

func TestAutorCrear(t *testing.T) {
	au := autor.Autor{
		Nombre: "Autos Test",
	}

	if err := au.Crear(); err != nil {
		t.Error(err)
		t.Fail()
	} else {

		id = au.Key
	}
}

func TestAutorListar(t *testing.T) {
	testCase := []struct {
		Nombre   string
		Res      int
		Consulta string
	}{
		{
			Nombre:   "Listar todos los datos",
			Res:      0,
			Consulta: "",
		},
		{
			Nombre:   "Filtar por nombre",
			Res:      1,
			Consulta: "Autos Test",
		},
	}

	for _, ts := range testCase {
		var au autor.Autor
		t.Run(ts.Nombre, func(t *testing.T) {
			if val, err := au.Listar(ts.Consulta); err != nil {
				t.Error(err)
				t.Fail()
			} else {
				if len(val) < ts.Res {
					t.Error("No se pudo hacer la cunsuta corecta")
					t.Fail()
				}
			}
		})

	}
}

func TestAutorVer(t *testing.T) {
	var au autor.Autor
	au.Ver(id)
	if au.Nombre == "" {
		t.Error("No se pudo traer el chofer")
		t.Fail()
	}
}

func TestAutorEditar(t *testing.T) {
	au := autor.Autor{Nombre: "Autos Test Editado"}
	if err := au.Editar(id); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestAutorEliminar(t *testing.T) {
	var au autor.Autor
	if err := au.Eliminar(id); err != nil {
		t.Error(err)
		t.Fail()
	}
}
