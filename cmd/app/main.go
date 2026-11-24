package main

import (
	controller "condominei/internal/controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main(){

	fmt.Println("Será que vai???");	

	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*.html")

	router.GET("/", controller.TextoSimples)
	router.GET("/:mensagem", controller.TextoDigitado)

    router.Run(":8002")

}