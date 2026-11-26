package usuario

import "github.com/gin-gonic/gin"

func Listar(c *gin.Context) {

	c.HTML(200, "index.html", gin.H{
		"mensagem": c.Param("mensagem"),
	})

}