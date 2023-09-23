package handlers

import "github.com/gin-gonic/gin"

func StaticHandler(c *gin.Context) {
	file := c.Param("file")
	c.File("static/" + file)
}
