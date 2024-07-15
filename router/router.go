package router

import "github.com/gin-gonic/gin"

func InicializadorDeRota() {
	r := gin.Default()
	setRotas(r)
	r.Run()
}
