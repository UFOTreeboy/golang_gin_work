package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index2.html", gin.H{
			"msg": "hello~ html",
		})
	})
	
	router.Run("127.0.0.1:8080")
}