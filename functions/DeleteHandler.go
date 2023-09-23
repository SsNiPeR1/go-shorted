package functions

import (
	"database/sql"

	"github.com/SsNiPeR1/go-shorted/config"
	"github.com/gin-gonic/gin"
)

func DeleteHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		admin_key := c.PostForm("admin_key")
		if admin_key != config.AdminKey {
			c.JSON(403, gin.H{
				"status": "error",
				"error":  "invalid admin key",
			})
			return
		}

		shorted := c.PostForm("shorted")
		if shorted == "" {
			c.JSON(400, gin.H{
				"status": "error",
				"error":  "shorted url not specified",
			})
			return
		}

		_, err := db.Exec("DELETE FROM urls WHERE shorted = $1", shorted)
		if err != nil {
			c.JSON(500, gin.H{
				"status": "error",
				"error":  "internal server error",
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "success",
			"data":   "url deleted",
		})
	}
}
