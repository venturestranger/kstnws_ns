package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

var httpClient http.Client = http.Client{}

func Log(data ...any) {
	log.Println("|", data)
}

func SendStatus(status int, c *gin.Context) {
	c.AbortWithStatus(status)
}

func SendString(value string, c *gin.Context) {
	c.String(200, value)
}

func SendData(jsonByte []Notification, c *gin.Context) {
	c.JSON(200, jsonByte)
}
