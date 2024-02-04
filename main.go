package main

import (
	"evento/search-api/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	path, exists := os.LookupEnv("PATH")
	if exists {
		fmt.Println(path)
	} else {
		return
	}

	r := gin.Default()
	r.Use(cors.Default())

	// search artist by name.
	r.GET("/search/artists/:term", func(c *gin.Context) {
		term := c.Param("term")
		data := handlers.GetArtistsByTerm(term)
		c.String(http.StatusOK, data)
	})

	// search events by artist mbid.
	r.GET("/search/events/:mbid", func(c *gin.Context) {
		mbid := c.Param("mbid")
		p := c.DefaultQuery("p", "1")
		data := handlers.GetEventsByArtistId(mbid, p)
		c.String(http.StatusOK, data)
	})

	// search events by artist mbid and year.
	r.GET("/search/events", func(c *gin.Context) {
		artistMbid := c.Query("artistMbid")
		year := c.Query("year") // shortcut for c.Request.URL.Query().Get("lastname")
		p := c.DefaultQuery("p", "1")

		data := handlers.GetArtistsEventsByYear(artistMbid, year, p)
		c.String(http.StatusOK, data)
	})

	r.Run(":8080")
}
