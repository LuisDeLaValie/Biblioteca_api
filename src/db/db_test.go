package db_test

import (
	"libreria/src/db"
	"testing"
)

var con db.Mongodb

func TestAbrirConexion(t *testing.T) {
	defer func() {
		con.Close()
		if r := recover(); r != nil {
			t.Error("Error al intetentar conexion")
			t.Error(r)
			t.Fail()
		}
	}()
	col := con.GetCollection("libros")

	if col == nil {
		t.Error("Error al abiri la conexion")
		t.Fail()
	}

}
