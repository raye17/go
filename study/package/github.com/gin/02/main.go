package main

import (
	"log"
	"net/http"

	"study/package/github.com/nicksnyder/i18n/01/pkg/common"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	common.InitBundle()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello,sxy!")
	})
	r.POST("/data", GetData)
	r.POST("/datas", GetDatas)
	r.Run(":8080")
}

type Data struct {
	Secret string `json:"secret"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}

func GetData(c *gin.Context) {
	var req Data
	//req.Secret = c.Query("secret")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	log.Println("req bind success")
	var data Data
	data = req
	//lang := c.GetHeader("Accept-Language")
	msg := common.Translate("", "UsernameExists", map[string]interface{}{
		"goods": "商品",
	})
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
		"data": data,
	})
}
func GetDatas(c *gin.Context) {
	var req Data
	//req.Secret = c.Query("secret")
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	log.Println("req bind success")
	var ds Data
	ds.Secret = c.Query("secret")
	if err := c.ShouldBindBodyWith(&ds, binding.JSON); err != nil {
		log.Println("ssss")
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"req":     req,
		"ds":      ds,
	})
}
