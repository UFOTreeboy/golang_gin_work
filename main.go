package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Printf("gogo %v \n", t)
	}
}
func main() {
	r := gin.Default()
	r.Use(LoggerHandler())
	//註冊一個Middleware

	r.GET("/", func(c *gin.Context) {
		fmt.Println("HELLO")
	})

	r.Run(":8000")
}