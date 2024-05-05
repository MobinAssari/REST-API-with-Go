package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEventHandler)
	err := server.Run(":8080")
	if err != nil {
		fmt.Print(err)
		return
	}
}

func getEventHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"massage": "hiii",
	})
}
