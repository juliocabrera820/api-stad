package controladores

import (
	"fmt"
	"github.com/julio-cabrera/api-stad/forma"
	"github.com/julio-cabrera/api-stad/modelos"
	"github.com/gin-gonic/gin"
)

var ModeloJugador = new(modelos.ModeloJugador)

type ControladorJugador struct{}

func (jugador *ControladorJugador) Agregar(c *gin.Context) {
	var data forma.CrearJugador
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"mensaje": "No valido", "forma": data})
		c.Abort()
		return
	}

	err := ModeloJugador.Agregar(data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Jugador no pudo ser agregado", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"mensaje": "Jugador Creado"})
}

func (jugador *ControladorJugador) Listar(c *gin.Context) {
	jugadores, err := ModeloJugador.Listar()
	if err != nil {
		c.JSON(404, gin.H{"message": "Error al obtener jugadores", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": jugadores})
	}
}

func (jugador *ControladorJugador) ObtenerJugador(c *gin.Context) {
	id := c.Param("id")
	jugadorEncontrado, err := ModeloJugador.ObtenerJugador(id)
	if err != nil {
		c.JSON(404, gin.H{"mensaje": "Jugador no encontrado", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": jugadorEncontrado})
	}
}

func (jugador *ControladorJugador) Actualizar(c *gin.Context) {
	id := c.Param("id")
	data := forma.ActualizarJugador{}

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"mensaje": "Parametros no validos"})
		c.Abort()
		return
	}
	err := ModeloJugador.Actualizar(id, data)
	if err != nil {
		c.JSON(406, gin.H{"mensaje": "Error al actualizar", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"mensaje": "Jugador Actualizado"})
}

func (jugador *ControladorJugador) Eliminar(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	err := ModeloJugador.Eliminar(id)
	if err != nil {
		c.JSON(406, gin.H{"mensaje": "Error al eliminar", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"mensaje": "Jugador Eliminado"})
}