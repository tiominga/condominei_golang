package routes

import (
	"condominei/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouts(router *gin.Engine){

	router.GET("/", controller.TextoSimples)
	router.GET("/:mensagem",controller.TextoDigitado)
	router.GET("/pessoa",controller.PessoaIndex)

}
