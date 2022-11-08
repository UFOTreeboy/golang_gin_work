package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
    "fmt"
    "log"

	_"github.com/mattn/go-sqlite3"
)

func main() {


    db, err := sql.Open("sqlite3", "test.db")

    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    sts := `
DROP TABLE IF EXISTS cars;
CREATE TABLE cars(id INTEGER PRIMARY KEY, name TEXT, price INT);
INSERT INTO cars(name, price) VALUES('Audi',52642);
INSERT INTO cars(name, price) VALUES('Mercedes',57127);
INSERT INTO cars(name, price) VALUES('Skoda',9000);
INSERT INTO cars(name, price) VALUES('Volvo',29000);
INSERT INTO cars(name, price) VALUES('Bentley',350000);
INSERT INTO cars(name, price) VALUES('Citroen',21000);
INSERT INTO cars(name, price) VALUES('Hummer',41400);
INSERT INTO cars(name, price) VALUES('Volkswagen',21600);
`
    _, err = db.Exec(sts)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("table cars created")




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
