package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	url := "localhost:8000"

	router := gin.New()
	setupRoutes(router)

	fmt.Println("Server online at:", url)
	router.Run(url)
}

func setupRoutes(router *gin.Engine) {
	router.GET("/ws/:id", func(c *gin.Context) {
		id := c.Param("id")
		addClient(c.Writer, c.Request, id)
	})
}
