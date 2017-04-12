package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"auto-ng/api"
	"fmt"
)

func main() {
	router := gin.Default()
	router.GET("/listAll", api.ListAll)
	router.GET("/list/:id", api.List)
	router.POST("/insert", api.Insert)
	router.POST("/update", api.Update)
	router.POST("/delete", api.Delete)
	router.GET("/etcdKeyList", api.QueryETCDKeyList)
	router.Run(":8003")
	fmt.Println("auto-ng is running at port : 8003 ...")
}