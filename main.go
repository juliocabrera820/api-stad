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
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "No Encontrado")
	})

	router.Run(":9090")
}