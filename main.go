package main

import (
	"example.com/REST-API/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEventHandler)
	server.POST("/events", CreateEvent)
	err := server.Run(":8080")
	if err != nil {
		fmt.Print(err)
		return
	}
}

func getEventHandler(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}
func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"message": "bad request"})
		return
	}
	event.Id = 1
	event.UserId = 1
	context.JSON(http.StatusCreated, event)
}
