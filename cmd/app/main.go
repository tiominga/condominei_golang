package main

import (
	"condominei/internal/routes"

	"github.com/gin-gonic/gin"
)

func main(){


	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*.html")

	routes.RegisterRouts(router)

	router.Run(":8002")

}