package forma

import (
	"gopkg.in/mgo.v2/bson"
)
type CrearEquipo struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Nombre string `json:"nombre" binding:"required"`
	Campeonatos int `json:"campeonatos" binding:"required"`
	Estadio string `json:"estadio" binding:"required"`
}

type ActualizarEquipo struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Nombre string `json:"nombre" binding:"required"`
	Campeonatos int `json:"campeonatos" binding:"required"`
	Estadio string `json:"estadio" binding:"required"`
}