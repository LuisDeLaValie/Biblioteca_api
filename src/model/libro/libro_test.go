package libro_test

import (
	"encoding/json"
	"libreria/src/model/autor"
	"libreria/src/model/libro"

	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var id primitive.ObjectID

var crearFlormulario libro.LibroFormulario
var editarFlormulario libro.LibroFormulario
var libroGeneral libro.Libro

func TestLibcroParseo(t *testing.T) {
	t.Run("formulario de crear", func(t *testing.T) {

		ida, _ := primitive.ObjectIDFromHex("63514e5fdcb2c4cbccd0b0ac")
		idb, _ := primitive.ObjectIDFromHex("63514e5fdcb2c4cbccd0b0ae")

		var formulario libro.LibroFormulario
		crearFlormulario = libro.LibroFormulario{
			Titulo:    "Libro Test",
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
		data := "{" +
			"\"titulo\": \"Libro Test\"," +
			"\"sinopsis\": \"Este es un libro de pruebas para Test\"," +
			"\"autores\": [" +
			"	\"63514e5fdcb2c4cbccd0b0ac\"," +
			"	\"63514e5fdcb2c4cbccd0b0ae\"" +
			"]," +
			"\"editorial\": \"Editoral de prubea\"," +
			"\"paginas\": 100," +
			"\"origen\": {" +
			"	\"nombre\": \"Test\"," +
			"	\"url\": \"www.TDTxLE/prueba/Test\"" +
			"}," +
			"\"Path\": \"1OWrmLzsFOMNPxq79-KdAW1s-ytXuw8ov\"" +
			"}"

		json.Unmarshal([]byte(data), &formulario)

		if res := crearFlormulario.Compare(formulario); res != nil {
			t.Error(*res)
			t.Fail()
		}

	})

	t.Run("formulario de editar", func(t *testing.T) {
		ida, _ := primitive.ObjectIDFromHex("63514e5fdcb2c4cbccd0b0ac")
		idb, _ := primitive.ObjectIDFromHex("63514e5fdcb2c4cbccd0b0ae")

		var formulario libro.LibroFormulario
		editarFlormulario = libro.LibroFormulario{
			Titulo:    "Libro Test Editado",
			Sipnosis:  "Este es un libro de pruebas para Test",
			Autores:   []primitive.ObjectID{ida, idb},
			Editorail: "Editoral de prubea",
			Pagina:    5,
			Origen: libro.LibroOrigen{
				Nombre: "Test",
				Url:    "www.TDTxLE/prueba/Test",
			},
			Path: "1OWrmLzsFOMNPxq79-KdAW1s-ytXuw8ov",
		}
		data := "{" +
			"\"titulo\": \"Libro Test Editado\"," +
			"\"sinopsis\": \"Este es un libro de pruebas para Test\"," +
			"\"autores\": [" +
			"	\"63514e5fdcb2c4cbccd0b0ac\"," +
			"	\"63514e5fdcb2c4cbccd0b0ae\"" +
			"]," +
			"\"editorial\": \"Editoral de prubea\"," +
			"\"pagina\": 5," +
			"\"origen\": {" +
			"	\"nombre\": \"Test\"," +
			"	\"url\": \"www.TDTxLE/prueba/Test\"" +
			"}," +
			"\"Path\": \"1OWrmLzsFOMNPxq79-KdAW1s-ytXuw8ov\"" +
			"}"

		json.Unmarshal([]byte(data), &formulario)

		if res := editarFlormulario.Compare(formulario); res != nil {
			t.Error(*res)
			t.Fail()
		}
	})

	t.Run("ver libro", func(t *testing.T) {
		ida, _ := primitive.ObjectIDFromHex("63514e5fdcb2c4cbccd0b0ac")
		idb, _ := primitive.ObjectIDFromHex("63514e5fdcb2c4cbccd0b0ae")

		libroGeneral = libro.Libro{
			Titulo:   "Libro Test",
			Sinopsis: "Este es un libro de pruebas para Test",
			Autores: []autor.Autor{
				{
					Key:    ida,
					Nombre: "Shirotaka",
				},
				{
					Key:    idb,
					Nombre: "Noboru Yamaguchi",
				},
			},
			Descargar:  "https://drive.google.com/uc?export=download&id=1OWrmLzsFOMNPxq79-KdAW1s-ytXuw8ov",
			Verr:       "https://drive.google.com/file/d/1OWrmLzsFOMNPxq79-KdAW1s-ytXuw8ov/view?usp=sharing",
			Paginacion: libro.LibroPaginacions{To: 0, End: 100},
			Origen: libro.LibroOrigen{
				Nombre: "Test",
				Url:    "www.TDTxLE/prueba/Test",
			},
		}
		var mylibro libro.Libro
		dato := "{" +
			"	\"titulo\": \"Libro Test\"," +
			"	\"Sinopsis\": \"Este es un libro de pruebas para Test\"," +
			"	\"autores\": [" +
			"		{" +
			"			\"key\": \"63514e5fdcb2c4cbccd0b0ac\"," +
			"			\"nombre\": \"Shirotaka\"" +
			"		}," +
			"		{" +
			"			\"key\": \"63514e5fdcb2c4cbccd0b0ae\"," +
			"			\"nombre\": \"Noboru Yamaguchi\"" +
			"		}" +
			"	]," +
			"	\"descargar\": \"https://drive.google.com/uc?export=download&id=1OWrmLzsFOMNPxq79-KdAW1s-ytXuw8ov\"," +
			"	\"ver\": \"https://drive.google.com/file/d/1OWrmLzsFOMNPxq79-KdAW1s-ytXuw8ov/view?usp=sharing\"," +
			"	\"paginacion\": {" +
			"		\"To\": 0," +
			"		\"End\": 100" +
			"	}," +
			"	\"origen\": {" +
			"		\"nombre\": \"Test\"," +
			"		\"url\": \"www.TDTxLE/prueba/Test\"" +
			"	}" +
			"}"
		json.Unmarshal([]byte(dato), &mylibro)

		if res := libroGeneral.Compare(mylibro); res != nil {
			t.Error(*res)
			t.Fail()
		}
	})
}

func TestLibroCrear(t *testing.T) {

	if val, err := crearFlormulario.Crear(); err != nil {
		t.Error(err)
		t.Fail()
	} else {
		id = val.Key
		if res := val.Compare(libroGeneral); res != nil {
			t.Error(*res)
			t.Fail()
		}
	}
}
func TestLibroListar(t *testing.T) {

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
			Consulta: "Libro Test",
		},
		{
			Nombre:   "Filtar por autor",
			Res:      1,
			Consulta: "Shirotaka",
		},
	}

	for _, ts := range testCase {
		var libro libro.Libro
		t.Run(ts.Nombre, func(t *testing.T) {
			if val, err := libro.Listar(ts.Consulta, false); err != nil {
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

func TestLibroVer(t *testing.T) {
	var mylibro libro.Libro

	if err := mylibro.Ver(id); err != nil {
		t.Error(err)
		t.Fail()
	} else {
		if res := mylibro.Compare(libroGeneral); res != nil {
			t.Error(*res)
			t.Fail()
		}
	}

}

func TestLibroEditar(t *testing.T) {

	if val, err := editarFlormulario.Editar(id); err != nil {
		t.Error(err)
		t.Fail()
	} else {
		libroaux := libroGeneral
		libroaux.Titulo = "Libro Test Editado"
		libroaux.Paginacion.To = 5
		if res := val.Compare(libroaux); res != nil {
			t.Error(*res)
			t.Fail()
		}
	}

}

func TestLibroeliminar(t *testing.T) {
	var libro libro.Libro

	if err := libro.Eliminar(&id, nil); err != nil {
		t.Error(err)
		t.Fail()
	}

}
