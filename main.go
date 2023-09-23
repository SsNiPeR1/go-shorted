package main

import (
	"fmt"
	"time"

	"database/sql"

	"github.com/SsNiPeR1/go-shorted/config"
	"github.com/SsNiPeR1/go-shorted/functions"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Server setup
	gin.SetMode(gin.ReleaseMode) // disable debug mode
	r := gin.Default()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Add the routes
	r.GET("/", functions.IndexHandler)
	r.GET("/static/:file", functions.StaticHandler)
	r.GET("/:shorted", functions.RedirectHandler(db)) // :shorted is a parameter which can be accessed with c.Param("shorted")

	r.POST("/db", functions.DBHandler(db)) // for debugging purposes only, can be used only by admins with the admin key POSTed
	r.POST("/vanity", functions.VanityHandler(db))
	r.POST("/delete", functions.DeleteHandler(db))
	r.POST("/shorten", functions.ShortenHandler(db))

	// log the time and date when the server starts, as well as the port it's listening on.
	fmt.Printf("[SHORTED] %s - %s | \033[32mListening on port 8080\033[0;0m\n", time.Now().Format("2006/01/02"), time.Now().Format("15:04:05"))
	r.Run(":8080")
}
