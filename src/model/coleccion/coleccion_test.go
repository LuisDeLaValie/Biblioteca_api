package coleccion_test

import (
	"libreria/src/model/coleccion"
	"libreria/src/model/libro"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var libroCreado libro.Libro
var coleccionCreado coleccion.Coleccion

func TestColeccionCrear(t *testing.T) {
	ida, _ := primitive.ObjectIDFromHex("63514e5fdcb2c4cbccd0b0ac")
	idb, _ := primitive.ObjectIDFromHex("63514e5fdcb2c4cbccd0b0ae")
	libroFormulario := libro.LibroFormulario{
		Titulo:    "Libro Test Coleccion",
		Sipnosis:  "Este es un libro de pruebas para Test",
		Autores:   []primitive.ObjectID{ida, idb},
		Editorail: "Editoral de prubea",
		Paginas:   100,
		Origen: libro.LibroOrigen{
			Nombre: "Test",
			Url:    "www.TDTxLE/prueba/Test",
		},
		Path: "1OWrmLzsFOMNPxq79-KdAW1s-ytXuw8ov",
	}

	if libroCreadoaux, err1 := libroFormulario.Crear(); err1 == nil {
		coleccionCreadoaux := coleccion.ColeccionFormulario{
			Titulo:   "Coleccion Test",
			Sinopsis: "Este es una coleccion de pruebas para Test",
			Libors:   []primitive.ObjectID{libroCreadoaux.Key},
		}

		if colaux, err2 := coleccionCreadoaux.Crear(); err2 == nil {
			coleccionCreado = *colaux
			libroCreado = *libroCreadoaux
			/*
				data := "{" +
					"\"titulo\": \"Coleccion Test\"," +
					"\"sinopsis\": \"Este es una coleccion de pruebas para Test\"," +
					"\"libros\": [" +
					"{" +
					"\"key\": \"638e8b67ed52434d1d66c1c4\"," +
					"\"titulo\": \"Libro Test Coleccion\"," +
					"\"Sinopsis\": \"Este es un libro de pruebas para Test\"," +
					"\"paginacion\": {" +
					"\"To\": 0," +
					"\"End\": 100" +
					"}," +
					"\"origen\": {" +
					"\"nombre\": \"Test\"," +
					"\"url\": \"www.TDTxLE/prueba/Test\"" +
					"}," +
					"}" +
					"]," +
					"}"
				var auxcreado coleccion.Coleccion
				json.Unmarshal([]byte(data), &auxcreado)

				if val := coleccionCreado.Compare(auxcreado); val != nil {
					t.Errorf(*val)
					t.Fail()
				} */
		} else {
			t.Error(err1)
			t.Fail()
		}
	} else {
		t.Error(err1)
		t.Fail()
	}
}

func TestColeccionListar(t *testing.T) {
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
			Consulta: "Coleccion Test",
		},
	}

	for _, ts := range testCase {
		var col coleccion.Coleccion
		t.Run(ts.Nombre, func(t *testing.T) {
			if val, err := col.Listar(ts.Consulta); err != nil {
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

func TestColeccionVer(t *testing.T) {
	var col coleccion.Coleccion
	if err := col.Ver(coleccionCreado.Key); err != nil {
		t.Error(err)
		t.Fail()
	} else {
		if val := col.Compare(coleccionCreado); val != nil {
			t.Error(err)
			t.Fail()
		}
	}
}

func TestColeccionEditar(t *testing.T) {

	coleccionCreadoaux := coleccion.ColeccionFormulario{
		Titulo:   "Coleccion Test Editado",
		Sinopsis: "Este es una coleccion de pruebas para Test",
		Libors:   []primitive.ObjectID{libroCreado.Key},
	}

	coleccionCreado.Titulo = "Coleccion Test Editado"
	if val, err := coleccionCreadoaux.Editar(coleccionCreado.Key); err != nil {
		t.Error(err)
		t.Fail()
	} else {
		if err := val.Compare(coleccionCreado); err != nil {
			t.Error(err)
			t.Fail()
		}
	}

}

func TestColeccionEliminar(t *testing.T) {
	var col coleccion.Coleccion
	if err := col.Eliminar(coleccionCreado.Key, true); err != nil {
		t.Error(err)
		t.Fail()
	}
}
