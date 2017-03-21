package api

import (
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
)

func TestFunc1(c *gin.Context) {
	// 回复一个200OK,在client的http-get的resp的body中获取数据
	c.String(http.StatusOK, "test1 OK")
}

// func2: 处理最基本的POST
func TestFunc2(c *gin.Context) {
	// 回复一个200 OK, 在client的http-post的resp的body中获取数据
	c.String(http.StatusOK, "test2 OK")
}
