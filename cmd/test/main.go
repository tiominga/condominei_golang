package main

import (
	"condominei/internal/model/usuario"
	"fmt"
)

func main(){

	obj_usuario := usuario.Usuario{}
	obj_usuario.Email = "tiominga@gmail.com"
	
	fmt.Println("Bem vindo " + obj_usuario.Email)



}


