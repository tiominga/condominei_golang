package controller

import (
	"fmt"

	"strings"

	"github.com/gin-gonic/gin"
)

func TextoSimples(c *gin.Context){

	fmt.Println("Oi mundo") // Terminal

    c.String(200, "Página funcionando!") // Navegador


} 

func TextoDigitado(c *gin.Context){

	mensagem := c.Param("mensagem")

	if strings.ContainsAny(mensagem, "0123456789") {
     c.String(400, "Não pode conter números")
     return
    }
	
	fmt.Println(mensagem)

	c.String(200, mensagem)


}



