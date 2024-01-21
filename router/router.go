package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	//inicializa o router utilizando as configurações Default do gin
	router := gin.Default()

	//Initialize Routes
	inicializaRoutes(router)
	// Rodando a API
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
