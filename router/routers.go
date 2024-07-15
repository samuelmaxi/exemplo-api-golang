package router

import (
	"api_exemplo/controllers"

	"github.com/gin-gonic/gin"
)

func setRotas(r *gin.Engine) {
	r.GET("/usuarios", controllers.ListaUsuarios)
	r.POST("/novousuario", controllers.NovoUsuario)
	r.PUT("/atualizar/:id", controllers.AtualizarUsuario)
	r.DELETE("/deletar/:id", controllers.DeletaUsuario)
}
