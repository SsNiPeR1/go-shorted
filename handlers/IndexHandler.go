package handlers

import (
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8") // needed for the browser to render the page
	c.File("templates/index.html")
}
