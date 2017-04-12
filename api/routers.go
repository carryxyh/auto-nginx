package api

import (
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
	"auto-ng/dao"
	"encoding/json"
	"strconv"
	"log"
	"auto-ng/etcd"
	"github.com/coreos/etcd/client"
	"auto-ng/config"
)

func ListAll(c *gin.Context) {

	//db操作
	var tmpls []*dao.Template
	tmpls = dao.QueryAll()
	b, err := json.Marshal(&tmpls)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, string(b))
}

func List(c *gin.Context) {

	//db操作
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
	c.String(http.StatusOK, string(b))
}

func Insert(c *gin.Context) {

	//db操作
	pjname := c.Param("pjname")
	content := c.Param("content")

	t := dao.Template{ProjectName:pjname, Content:content}
	id := dao.Insert(t)
	c.String(http.StatusOK, string(id))

	//etcd操作
	etcd.UpdateKey(pjname, content)
}

func Update(c *gin.Context) {

	//db操作
	id := c.Param("id")
	iId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	pjname := c.Param("pjname")
	content := c.Param("content")

	t := dao.Template{Id:iId, ProjectName:pjname, Content:content}
	b := dao.Update(t)
	if b {
		c.String(http.StatusOK, "insert ok !")
	} else {
		c.String(http.StatusInternalServerError, "insert failed !")
	}

	//etcd操作
	etcd.UpdateKey(pjname, content)
}

func Delete(c *gin.Context) {

	//db操作
	id := c.Param("id")
	iId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	b := dao.DeleteById(iId)
	if b {
		c.String(http.StatusOK, "insert ok !")
	} else {
		c.String(http.StatusInternalServerError, "insert failed !")
	}

	//etcd操作
	pjname := c.Param("pjname")
	opts := client.DeleteOptions{}
	etcd.DeleteKey(pjname, &opts)
}

func QueryETCDKeyList(c *gin.Context) {
	conf := config.GetETCDconfig()
	response, err := etcd.QueryKeyList(conf.BasePath)
	if err != nil {
		log.Fatal(err)
	}
	n := response.Node
	b, errJ := json.Marshal(n)
	if errJ != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, string(b))
}