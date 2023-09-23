package handlers

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/SsNiPeR1/go-shorted/config"
	"github.com/gin-gonic/gin"
	"github.com/noirbizarre/gonja"
)

func ShortenHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8") // needed for the browser to render the page

		// get the URL that should be shortened
		url := c.PostForm("url")
		time := int32(time.Now().Unix())
		match := MatchURL(url)

		tpl500, _ := gonja.FromFile("templates/500.html")
		rendered500, _ := tpl500.Execute(gonja.Context{"ginversion": gin.Version})

		tpl, err := gonja.FromFile("templates/shorted.html")
		if err != nil {
			c.String(500, rendered500)
			return
		}

		if !match {
			c.File("templates/invalid.html") // tell the user that the entered URL is invalid
			return
		}

		// check if there is already this url in the database
		var shorted string
		err = db.QueryRow("SELECT shorted FROM urls WHERE url = $1", url).Scan(&shorted)
		if err == nil {
			fullurl := config.Website + shorted
			rendered, _ := tpl.Execute(gonja.Context{"url": fullurl})
			c.String(200, rendered)

			return
		}

		// generate a random string while checking if it already exists in the database
		for {
			result := make([]byte, 6) // we allocate a slice of bytes with the length of 6

			for i := range result {
				result[i] = config.Charset[rand.Intn(len(config.Charset))] // add a random char to the slice
			}
			shorted = string(result) // convert the slice to a string

			err := db.QueryRow("SELECT shorted FROM urls WHERE shorted = $1", shorted).Scan(&shorted)
			if err != nil {
				break // means that the shorted URL doesn't exist in the database, therefore can be used
			}
		}

		// render the template
		fullurl := config.Website + shorted // to render the full URL instead of just path
		rendered, _ := tpl.Execute(gonja.Context{"url": fullurl})

		// insert the url into the database
		_, err = db.Exec("INSERT INTO urls (url, shorted, created_at) VALUES ($1, $2, $3)", url, shorted, time)
		if err != nil {
			c.String(500, rendered500)
		}
		c.String(200, rendered)
	}
}
