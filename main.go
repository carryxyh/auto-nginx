package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"auto-ng/api"
)

func main() {
	router := gin.Default()
	router.GET("/list", api.TestFunc1)
	router.POST("/insert", api.TestFunc2)
	router.Run(":8003")
}