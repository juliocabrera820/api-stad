package modelos


import (
	"github.com/julio-cabrera/api-stad/bd"
	"github.com/julio-cabrera/api-stad/forma"
	"gopkg.in/mgo.v2/bson"
)

type Equipo struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Nombre string `json:"nombre"`
	Campeonatos int `json:"campeonatos"`
	Estadio string `json:"estadio"`
}

type ModeloEquipo struct{}

var servidor = "127.0.0.1"

var bdConectada = bd.NuevaConexion(servidor)

func (m *ModeloEquipo) Agregar(data forma.CrearEquipo) error {
	coleccion := bdConectada.Usar("stad", "equipos")
	err := coleccion.Insert(bson.M{"nombre": data.Nombre, "campeonatos": data.Campeonatos, "estadio": data.Estadio})
	return err
}

func (m *ModeloEquipo) Listar() (equipos []Equipo, err error) {
	coleccion := bdConectada.Usar("stad", "equipos")
	err = coleccion.Find(bson.M{}).All(&equipos)
	return equipos, err
}

func (m *ModeloEquipo) ObtenerEquipo(id string) (equipo Equipo, err error) {
	coleccion := bdConectada.Usar("stad", "equipos")
	err = coleccion.FindId(bson.ObjectIdHex(id)).One(&equipo)
	return equipo, err
}

func (m *ModeloEquipo) Actualizar(id string, data forma.ActualizarEquipo) (err error) {
	coleccion := bdConectada.Usar("stad", "equipos")
	err = coleccion.UpdateId(bson.ObjectIdHex(id), data)
	return err
}

func (m *ModeloEquipo) Eliminar(id string) (err error) {
	coleccion := bdConectada.Usar("stad", "equipos")
	err = coleccion.RemoveId(bson.ObjectIdHex(id))
	return err
}