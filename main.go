package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Backend")
	})
	r.Run(":8080")
	log.Println("Backend is Running")
}
