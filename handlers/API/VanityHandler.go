package api

import (
	"database/sql"
	"time"

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

		time := int32(time.Now().Unix())
		url := c.PostForm("url")
		shorted := c.PostForm("shorted")

		// check if the shorted URL is already in the database
		var shorted_check string
		err := db.QueryRow("SELECT shorted FROM urls WHERE shorted = $1", shorted).Scan(&shorted_check)
		if err == nil {
			c.JSON(403, gin.H{
				"status": "error",
				"error":  "shorted URL already exists",
			})
			return
		}

		_, err = db.Exec("INSERT INTO urls (url, shorted, created_at) VALUES ($1, $2, $3)", url, shorted, time)
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
