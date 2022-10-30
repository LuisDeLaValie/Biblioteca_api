package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Autor struct {
	Key    primitive.ObjectID `bson:"_id,omitempty" json:"key,omitempty"`
	Nombre string             `bson:"nombre" json:"nombre"`
}
