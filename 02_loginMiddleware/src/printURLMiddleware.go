package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func printURLMiddleware(c *gin.Context) {

	fmt.Println("Print Middleware: ", c.Request.URL.Path)

	c.Next()
}
