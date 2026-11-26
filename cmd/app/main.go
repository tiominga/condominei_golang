package main

import (
	"condominei/internal/home"
	"condominei/internal/usuario"

	"github.com/gin-gonic/gin"
)

func main(){


	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*.html")

	home.RegisterRoutes(router)
	usuario.RegisterRoutes(router)

	router.Run(":8002")

}