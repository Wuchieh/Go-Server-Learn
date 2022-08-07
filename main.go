package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Panicln("Error:", err.Error())
	}

	r.GET("/status", status)
	r.GET("/", index)

	err = r.Run("127.0.0.1:8080")
	if err != nil {
		log.Panicln("Error:", err.Error())
	}
}

func index(c *gin.Context) {
	fmt.Println("index")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "載入完成",
	})
}

func status(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "200",
		"message": "server Online",
	})
}
