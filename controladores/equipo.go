package controladores

import (
	"fmt"
	"github.com/julio-cabrera/api-stad/forma"
	"github.com/julio-cabrera/api-stad/modelos"
	"github.com/gin-gonic/gin"
)

var ModeloEquipo = new(modelos.ModeloEquipo)

type ControladorEquipo struct{}

func (equipo *ControladorEquipo) Agregar(c *gin.Context) {
	var data forma.CrearEquipo
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"mensaje": "No valido", "forma": data})
		c.Abort()
		return
	}

	err := ModeloEquipo.Agregar(data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Equipo no pudo ser agregado", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"mensaje": "Equipo Creado"})
}

func (equipo *ControladorEquipo) Listar(c *gin.Context) {
	equipos, err := ModeloEquipo.Listar()
	if err != nil {
		c.JSON(404, gin.H{"message": "Error al obtener equipos", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": equipos})
	}
}

func (equipo *ControladorEquipo) ObtenerEquipo(c *gin.Context) {
	id := c.Param("id")
	equipoEncontrado, err := ModeloEquipo.ObtenerEquipo(id)
	if err != nil {
		c.JSON(404, gin.H{"mensaje": "Equipo no encontrado", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": equipoEncontrado})
	}
}

func (equipo *ControladorEquipo) Actualizar(c *gin.Context) {
	id := c.Param("id")
	data := forma.ActualizarEquipo{}

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"mensaje": "Parametros no validos"})
		c.Abort()
		return
	}
	err := ModeloEquipo.Actualizar(id, data)
	if err != nil {
		c.JSON(406, gin.H{"mensaje": "Error al actualizar", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"mensaje": "Equipo Actualizado"})
}

func (equipo *ControladorEquipo) Eliminar(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	err := ModeloEquipo.Eliminar(id)
	if err != nil {
		c.JSON(406, gin.H{"mensaje": "Error al eliminar", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"mensaje": "Equipo Eliminado"})
}