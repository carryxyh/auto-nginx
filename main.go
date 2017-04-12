package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"auto-ng/api"
)

func main() {
	router := gin.Default()
	router.GET("/listAll", api.ListAll)
	router.GET("/list/:id", api.List)
	router.POST("/insert", api.Insert)
	router.POST("/update", api.Update)
	router.POST("/delete", api.Delete)
	router.Run(":8003")
}