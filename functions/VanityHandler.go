package functions

import (
	"database/sql"

	"github.com/SsNiPeR1/go-shorted/config"
	"github.com/gin-gonic/gin"
)

func VanityHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		admin_key := c.PostForm("admin_key")
		if admin_key != config.AdminKey {
			c.JSON(403, gin.H{
				"status": "error",
				"error":  "invalid admin key",
			})
			return
		}

		url := c.PostForm("url")
		shorted := c.PostForm("shorted")

		_, err := db.Exec("INSERT INTO urls (url, shorted) VALUES ($1, $2)", url, shorted)
		if err != nil {
			c.JSON(500, gin.H{
				"status": "error",
				"error":  "internal server error",
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "success",
			"data":   "url added",
			"url":    config.Website + shorted,
		})
	}
}
