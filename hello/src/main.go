package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("D:/goworkspace/src/hello/templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	sample := r.Group("/html")
	sample.GET("/sample", func(c *gin.Context) {
		c.HTML(http.StatusOK, "html-sample2.html", nil)
	})

	r.Static("/public", "D:/goworkspace/src/hello/public")

	r.Run(":3000")
}
