package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"time"
)

var (
	r = gin.Default()
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {

	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	// 載入順序很重要 若先載入template會因為沒有FuncMap報錯
	r.LoadHTMLGlob("./templates/*")

	r.GET("/status", status)
	r.GET("/", index)

	err := r.Run(":8080")
	if err != nil {
		log.Panicln("Error:", err.Error())
	}
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"message": "載入完成",
		"now":     time.Now(),
	})
}

func status(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "200",
		"message": "server Online",
	})
}
