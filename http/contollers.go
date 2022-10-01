package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}

func controllers(router *gin.RouterGroup) {
	groupControllers := router.Group("/controllers")
	groupControllers.GET("/hw", GetHelloWorld)
}
