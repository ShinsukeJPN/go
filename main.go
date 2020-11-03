package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"qiita/controller"
	"qiita/middleware"
)

func main() {
    engine:= gin.Default()

    engine.LoadHTMLGlob("templates/*")
		engine.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "message": "hello world",
        })
    })
    engine.Run(":5000")
}