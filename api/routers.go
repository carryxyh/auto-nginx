package api

import (
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
	"auto-ng/dao"
	"encoding/json"
	"strconv"
	"log"
)

func ListAll(c *gin.Context) {
	var tmpls []*dao.Template
	tmpls = dao.QueryAll()
	b, err := json.Marshal(&tmpls)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, string(b))
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
	c.String(http.StatusOK, string(b))
}

func Insert(c *gin.Context) {
	pjname := c.Param("pjname")
	content := c.Param("content")

	t := dao.Template{ProjectName:pjname, Content:content}
	id := dao.Insert(t)
	c.String(http.StatusOK, string(id))
}

func Update(c *gin.Context) {
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
}

func Delete(c *gin.Context) {
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
}