package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// inicia o roteador Gin padrão
	router := gin.Default()

	// Rotas
	router.GET("/ping", getPing)
	router.POST("/post", postMessage)
	router.PUT("/put", putMessage)
	router.DELETE("/delete", deleteMessage)

	// Inicia um servidor na porta 8080
	err := router.Run(":8080")
	if err != nil {
		panic("Falha ao iniciar o servidor")
	}
}

// Função para a rota GET /ping
func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// Função para a rota POST /post
func postMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Requisição POST recebida"})
}

// Função para a rota PUT /put
func putMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Requisição PUT recebida"})
}

// Função para a rota DELETE /delete
func deleteMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Requisição DELETE recebida"})
}
