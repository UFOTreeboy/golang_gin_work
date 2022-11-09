package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {


	router := gin.Default()
	router.LoadHTMLGlob("template/*")//讀取template的html


	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", gin.H{})
	})//註冊路由

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "page.html", gin.H{
			"msg": "hello~ html",
		})
	})
	router.Run("127.0.0.1:8080")

}
