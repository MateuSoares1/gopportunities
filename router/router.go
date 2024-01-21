package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//inicializa o router utilizando as configurações Default do gin
	router := gin.Default()
	// Definindo uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Rodando a API
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}