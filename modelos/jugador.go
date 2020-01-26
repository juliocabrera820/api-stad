package modelos

import (
	"github.com/julio-cabrera/api-stad/forma"
	"gopkg.in/mgo.v2/bson"
)

type Jugador struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Nombre string `json:"nombre"`
	Edad int `json:"edad"`
	Posicion string `json:"posicion"`
	Numero int `json:"numero"`
	UltimoEquipo string `json:"ultimoEquipo"`
}

type ModeloJugador struct{}

func (m *ModeloJugador) Agregar(data forma.CrearJugador) error {
	coleccion := bdConectada.Usar("stad", "jugadores")
	err := coleccion.Insert(bson.M{"nombre": data.Nombre, "edad": data.Edad, "posicion": data.Posicion, "numero": data.Numero, "ultimoEquipo":data.UltimoEquipo})
	return err
}

func (m *ModeloJugador) Listar() (jugadores []Jugador, err error) {
	coleccion := bdConectada.Usar("stad", "jugadores")
	err = coleccion.Find(bson.M{}).All(&jugadores)
	return jugadores, err
}

func (m *ModeloJugador) ObtenerJugador(id string) (jugador Jugador, err error) {
	coleccion := bdConectada.Usar("stad", "jugadores")
	err = coleccion.FindId(bson.ObjectIdHex(id)).One(&jugador)
	return jugador, err
}

func (m *ModeloJugador) Actualizar(id string, data forma.ActualizarJugador) (err error) {
	coleccion := bdConectada.Usar("stad", "jugadores")
	err = coleccion.UpdateId(bson.ObjectIdHex(id), data)
	return err
}

func (m *ModeloJugador) Eliminar(id string) (err error) {
	coleccion := bdConectada.Usar("stad", "jugadores")
	err = coleccion.RemoveId(bson.ObjectIdHex(id))
	return err
}