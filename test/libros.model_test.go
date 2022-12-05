package test

import (
	"encoding/json"
	"libreria/src/model"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var id primitive.ObjectID

func TestLibroCrear(t *testing.T) {
	var formulario model.LibroFormulario
	dataformulario := "{" +
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

	t.Run("comprobar parseo de datos", func(t *testing.T) {
		json.Unmarshal([]byte(dataformulario), &formulario)

		ida, _ := primitive.ObjectIDFromHex("63514e5fdcb2c4cbccd0b0ac")
		idb, _ := primitive.ObjectIDFromHex("63514e5fdcb2c4cbccd0b0ae")

		auxformulario := model.LibroFormulario{
			Titulo:    "Libro Test",
			Sipnosis:  "Este es un libro de pruebas para Test",
			Autores:   []primitive.ObjectID{ida, idb},
			Editorail: "Editoral de prubea",
			Paginas:   100,
			Origen:    model.LibroOrigen{Nombre: "Test", Url: "www.TDTxLE/prueba/Test"},
			Path:      "1OWrmLzsFOMNPxq79-KdAW1s-ytXuw8ov",
		}

		if formulario.Titulo != auxformulario.Titulo {
			t.Log(formulario.Titulo)
			t.Error("No se pudeo parsear el titulo")
			t.Fail()
		}
		if formulario.Sipnosis != auxformulario.Sipnosis {
			t.Log(formulario.Sipnosis)
			t.Error("No se pudeo parsear el Sipnosis")
			t.Fail()
		}
		// if formulario.Autores != auxformulario.Autores {
		// t.Log(formulario.Autores)
		// 	t.Error("No se pudeo parsear el Autores")
		// 	t.Fail()
		// }
		if formulario.Editorail != auxformulario.Editorail {
			t.Log(formulario.Editorail)
			t.Error("No se pudeo parsear el Editorail")
			t.Fail()
		}
		if formulario.Paginas != auxformulario.Paginas {
			t.Log(formulario.Paginas)
			t.Error("No se pudeo parsear el Paginas")
			t.Fail()
		}
		if formulario.Origen != auxformulario.Origen {
			t.Log(formulario.Origen)
			t.Error("No se pudeo parsear el Origen")
			t.Fail()
		}
		if formulario.Path != auxformulario.Path {
			t.Log(formulario.Path)
			t.Error("No se pudeo parsear el Path")
			t.Fail()
		}
	})

	t.Run("Crear libro", func(t *testing.T) {
		if val, err := formulario.Crear(); err != nil {
			t.Error(err)
			t.Fail()
		} else {
			id = val.Key
			var libro model.Libro
			datalibro := "{" +
				"	\"key\": \"638d4a00ff85bd61e20e2e00\"," +
				"	\"titulo\": \"Libro Test\"," +
				"	\"sipnosis\": \"Este es un libro de pruebas para Test\"," +
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
				"	\"paginacion\": {" +
				"		\"To\": 0," +
				"		\"End\": 100" +
				"	}," +
				"	\"origen\": {" +
				"		\"nombre\": \"Test\"," +
				"		\"url\": \"www.TDTxLE/prueba/Test\"" +
				"	}," +
				"	\"creado\": \"2022-12-05T01:31:44.916Z\"" +
				"}"

			json.Unmarshal([]byte(datalibro), &libro)

			if val.Titulo != libro.Titulo {
				t.Errorf("Titulo no coincide ::\n A) %s\nB) %s", val.Titulo, libro.Titulo)
				t.Fail()
			}
			if val.Sinopsis != libro.Sinopsis {
				t.Errorf("Sinopsis no coincide ::\n A) %s\nB) %s", val.Sinopsis, libro.Sinopsis)
				t.Fail()
			}
			if len(val.Autores) != len(libro.Autores) {
				t.Errorf("Autores no coincide %d", len(val.Autores))
				t.Fail()
			}
			if val.Editorail != libro.Editorail {
				t.Errorf("Editorail no coincide ::\n A) %s\nB) %s", val.Editorail, libro.Editorail)
				t.Fail()
			}
			/* if val.Descargar != libro.Descargar {
				t.Errorf("Descargar no coincide ::\n A) %s\nB) %s", val.Descargar, libro.Descargar)
				t.Fail()
			}
			if val.Path != libro.Path {
				t.Errorf("Path no coincide ::\n A) %s\nB) %s", val.Path, libro.Path)
				t.Fail()
			}
			if val.Verr != libro.Verr {
				t.Errorf("Ver no coincide ::\n A) %s\nB) %s", val.Verr, libro.Verr)
				t.Fail()
			} */
			if val.Paginacion != libro.Paginacion {
				t.Errorf("Paginacion no coincide ::\n {to:%d, end:%d}\n{to:%d, end:%d}", val.Paginacion.To, val.Paginacion.End, libro.Paginacion.To, libro.Paginacion.End)
				t.Fail()
			}
			if val.Origen != libro.Origen {
				t.Errorf("Origen no coincide ::\n A) %s\nB) %s", val.Origen, libro.Origen)
				t.Fail()
			}

		}
	})
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
		var libro model.Libro
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
	var libro model.Libro

	if err := libro.Ver(id); err != nil {
		t.Error(err)
		t.Fail()
	} else {
		data := "{" +
			"	\"titulo\": \"Libro Test\"," +
			"	\"sipnosis\": \"Este es un libro de pruebas para Test\"," +
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
		var auxlibro model.Libro

		json.Unmarshal([]byte(data), &auxlibro)
		if libro.Titulo != auxlibro.Titulo {
			t.Errorf("Titulo no coincide ::\n A) %s\nB) %s", libro.Titulo, auxlibro.Titulo)
			t.Fail()
		}
		if libro.Sinopsis != auxlibro.Sinopsis {
			t.Errorf("Sinopsis no coincide ::\n A) %s\nB) %s", libro.Sinopsis, auxlibro.Sinopsis)
			t.Fail()
		}
		if len(libro.Autores) != len(auxlibro.Autores) {
			t.Errorf("Autores no coincide %d", len(libro.Autores))
			t.Fail()
		}
		if libro.Editorail != auxlibro.Editorail {
			t.Errorf("Editorail no coincide ::\n A) %s\nB) %s", libro.Editorail, auxlibro.Editorail)
			t.Fail()
		}
		if libro.Descargar != auxlibro.Descargar {
			t.Errorf("Descargar no coincide ::\n A) %s\nB) %s", libro.Descargar, auxlibro.Descargar)
			t.Fail()
		}
		/* if libro.Path != auxlibro.Path {
			t.Errorf("Path no coincide ::\n A) %s\nB) %s", libro.Path, auxlibro.Path)
			t.Fail()
		} */
		if libro.Verr != auxlibro.Verr {
			t.Errorf("Ver no coincide ::\n A) %s\nB) %s", libro.Verr, auxlibro.Verr)
			t.Fail()
		}
		if libro.Paginacion != auxlibro.Paginacion {
			t.Errorf("Paginacion no coincide ::\n {to:%d, end:%d}\n{to:%d, end:%d}", libro.Paginacion.To, libro.Paginacion.End, auxlibro.Paginacion.To, auxlibro.Paginacion.End)
			t.Fail()
		}
		if libro.Origen != auxlibro.Origen {
			t.Errorf("Origen no coincide ::\n A) %s\nB) %s", libro.Origen, auxlibro.Origen)
			t.Fail()
		}

	}

}

func TestLibroEditar(t *testing.T) {

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

	var libro model.LibroFormulario
	json.Unmarshal([]byte(data), &libro)

	if val, err := libro.Editar(id); err != nil {
		t.Error(err)
		t.Fail()
	} else {
		var libro model.Libro
		datalibro := "{" +
			"	\"titulo\": \"Libro Test Editado\"," +
			"	\"sipnosis\": \"Este es un libro de pruebas para Test\"," +
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
			"	\"paginacion\": {" +
			"		\"To\": 5," +
			"		\"End\": 100" +
			"	}," +
			"	\"origen\": {" +
			"		\"nombre\": \"Test\"," +
			"		\"url\": \"www.TDTxLE/prueba/Test\"" +
			"	}," +
			"	\"creado\": \"2022-12-05T01:31:44.916Z\"" +
			"}"

		json.Unmarshal([]byte(datalibro), &libro)

		if val.Titulo != libro.Titulo {
			t.Errorf("Titulo no coincide ::\n A) %s\nB) %s", val.Titulo, libro.Titulo)
			t.Fail()
		}
		if val.Sinopsis != libro.Sinopsis {
			t.Errorf("Sinopsis no coincide ::\n A) %s\nB) %s", val.Sinopsis, libro.Sinopsis)
			t.Fail()
		}
		if len(val.Autores) != len(libro.Autores) {
			t.Errorf("Autores no coincide %d", len(val.Autores))
			t.Fail()
		}
		if val.Editorail != libro.Editorail {
			t.Errorf("Editorail no coincide ::\n A) %s\nB) %s", val.Editorail, libro.Editorail)
			t.Fail()
		}
		/* if val.Descargar != libro.Descargar {
			t.Errorf("Descargar no coincide ::\n A) %s\nB) %s", val.Descargar, libro.Descargar)
			t.Fail()
		}
		if val.Path != libro.Path {
			t.Errorf("Path no coincide ::\n A) %s\nB) %s", val.Path, libro.Path)
			t.Fail()
		}
		if val.Verr != libro.Verr {
			t.Errorf("Ver no coincide ::\n A) %s\nB) %s", val.Verr, libro.Verr)
			t.Fail()
		} */
		if val.Paginacion != libro.Paginacion {
			t.Errorf("Paginacion no coincide ::\n {to:%d, end:%d}\n{to:%d, end:%d}", val.Paginacion.To, val.Paginacion.End, libro.Paginacion.To, libro.Paginacion.End)
			t.Fail()
		}
		if val.Origen != libro.Origen {
			t.Errorf("Origen no coincide ::\n A) %s\nB) %s", val.Origen, libro.Origen)
			t.Fail()
		}
	}

}

func TestLibroeliminar(t *testing.T) {
	var libro model.Libro

	if err := libro.Eliminar(id); err != nil {
		t.Error(err)
		t.Fail()
	}

}
