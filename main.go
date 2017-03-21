package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"auto-ng/api"
)

func main() {
	router := gin.Default()
	router.GET("/test1", api.TestFunc1)
	router.POST("/test2", api.TestFunc2)
	router.Run(":8003")
}