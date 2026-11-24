package controller

import "github.com/gin-gonic/gin"

func PessoaIndex(c *gin.Context) {

	c.HTML(200, "index.html", gin.H{
		"mensagem": c.Param("mensagem"),
	})

}