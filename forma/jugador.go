package forma

import (
	"gopkg.in/mgo.v2/bson"
)
type CrearJugador struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Nombre string `json:"nombre" binding:"required"`
	Edad int `json:"edad" binding:"required"`
	Posicion string `json:"posicion" binding:"required"`
	Numero int `json:"numero" binding:"required"`
	UltimoEquipo string `json:"ultimoEquipo" binding:"required"`
}

type ActualizarJugador struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Nombre string `json:"nombre" binding:"required"`
	Edad int `json:"edad" binding:"required"`
	Posicion string `json:"posicion" binding:"required"`
	Numero int `json:"numero" binding:"required"`
	UltimoEquipo string `json:"ultimoEquipo" binding:"required"`
}