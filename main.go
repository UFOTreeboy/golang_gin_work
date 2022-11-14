package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
)

type book struct{
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Quantity int   `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c*gin.Context){
	c.IndentedJSON(http.StatusOK, books)

}


func main(){
	router := gin.Default()
	router.GET("/books",getBooks)
	router.Run("localhost:8080")


}


/*
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
*/
