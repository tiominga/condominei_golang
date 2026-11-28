package home

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine){
	
	router.GET("/",Home)
	
}
