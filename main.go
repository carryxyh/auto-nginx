package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func testFunc1(c *gin.Context) {
	// 回复一个200OK,在client的http-get的resp的body中获取数据
	c.String(http.StatusOK, "test1 OK")
}

// func2: 处理最基本的POST
func testFunc2(c *gin.Context) {
	// 回复一个200 OK, 在client的http-post的resp的body中获取数据
	c.String(http.StatusOK, "test2 OK")
}

func main() {
	router := gin.Default()
	router.GET("/test1", testFunc1)
	router.POST("/test2", testFunc2)
	router.Run(":8003")
}