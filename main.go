package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/julio-cabrera/api-stad/controladores"
)

func main() {
	router := gin.Default()

	{
		jugador := new(controladores.ControladorJugador)
		router.POST("/jugadores", jugador.Agregar)
		router.GET("/jugadores", jugador.Listar)
		router.GET("/jugadores/:id", jugador.ObtenerJugador)
		router.PUT("/jugadores/:id", jugador.Actualizar)
		router.DELETE("/jugadores/:id", jugador.Eliminar)
		equipo := new (controladores.ControladorEquipo)
		router.POST("/equipos", equipo.Agregar)
		router.GET("/equipos", equipo.Listar)
		router.GET("/equipos/:id", equipo.ObtenerEquipo)
		router.PUT("/equipos/:id", equipo.Actualizar)
		router.DELETE("/equipos/:id", equipo.Eliminar)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "No Encontrado")
	})

	router.Run(":9090")
}