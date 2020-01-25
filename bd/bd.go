package bd

import (
	"gopkg.in/mgo.v2"
)

type Conexion struct {
	sesion *mgo.Session
}

func NuevaConexion(host string) (conn *Conexion) {
	sesion, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	sesion.SetMode(mgo.Monotonic, true)
	conn = &Conexion{sesion}
	return conn
}

func (conn *Conexion) Usar(nombreBD, coleccion string) (collection *mgo.Collection) {
	return conn.sesion.DB(nombreBD).C(coleccion)
}

func (conn *Conexion) Cerrar() {
	conn.sesion.Close()
	return
}