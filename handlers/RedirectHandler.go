package handlers

import (
	"database/sql"
	"strings"

	"github.com/gin-gonic/gin"
)

func RedirectHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the shorted url from the request
		shorted := c.Param("shorted")

		// Query the database for the shorted url
		var url string
		err := db.QueryRow("SELECT url FROM urls WHERE shorted = $1", shorted).Scan(&url)
		if err != nil {
			c.JSON(404, gin.H{
				"status": "error",
				"error":  "url not found",
			})
			return
		}

		// add http:// in case the url doesn't have any protocol specified
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}

		// Redirect the user to the url
		c.Redirect(302, url)
	}
}
