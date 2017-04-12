package api

import (
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
	"auto-ng/dao"
	"encoding/json"
	"strconv"
	"log"
)

//func TestFunc1(c *gin.Context) {
//	// 回复一个200OK,在client的http-get的resp的body中获取数据
//	c.String(http.StatusOK, "test1 OK")
//}
//
//// func2: 处理最基本的POST
//func TestFunc2(c *gin.Context) {
//	// 回复一个200 OK, 在client的http-post的resp的body中获取数据
//	c.String(http.StatusOK, "test2 OK")
//}

func ListAll(c *gin.Context) {

}

func List(c *gin.Context) {
	id := c.Param("id")
	iId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	tmpl := dao.QueryById(iId)
	b, errJ := json.Marshal(tmpl)
	if errJ != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, b)
}

func Insert(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}