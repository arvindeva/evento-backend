package main

import (
	"evento/search-api/handlers"
	"fmt"
	"log"
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
	r.GET("/search/artists/:term", handlers.GetArtistsByTerm)
	// search events by artist mbid.
	r.GET("/search/events/:mbid", handlers.GetEventsByArtistId)
	// search events by artist mbid and year.
	r.GET("/search/events", handlers.GetArtistsEventsByYear)
	// search event by event id
	r.GET("/search/event/:event_id", handlers.GetEventById)

	r.Run(":8080")
}
