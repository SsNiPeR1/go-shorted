package api

import (
	"database/sql"

	"github.com/SsNiPeR1/go-shorted/config"
	"github.com/SsNiPeR1/go-shorted/structs"
	"github.com/gin-gonic/gin"
)

func DBHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the admin key from the request
		admin_key := c.PostForm("admin_key")
		if admin_key != config.AdminKey {
			c.JSON(403, gin.H{
				"status": "error",
				"error":  "invalid admin key",
			})
			return
		}

		ret := []structs.Urls{}

		var (
			url        string
			shorted    string
			created_at int32
		)

		rows, err := db.Query("SELECT * FROM urls")
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&url, &shorted, &created_at)
			if err != nil {
				panic(err) // sorry
			}
			ret = append(ret, structs.Urls{URL: url, Shorted: shorted, Created_at: created_at})
		}

		c.JSON(200, gin.H{
			"status": "success",
			"data":   ret,
		})
	}
}
