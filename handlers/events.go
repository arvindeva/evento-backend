package handlers

import (
	"evento/search-api/utils"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetEventsByArtistId(c *gin.Context) {
	fmt.Println("Making GET request...")
	fmt.Println("Searching events...")
	setlistFmApiKey, _ := os.LookupEnv("SETLIST_FM_API_KEY")
	mbid := c.Param("mbid")
	p := c.DefaultQuery("p", "1")

	// make GET request to API to get user by ID
	apiUrl := fmt.Sprintf("https://api.setlist.fm/rest/1.0/artist/%s/setlists?p=%s", mbid, p)
	request, error := http.NewRequest("GET", apiUrl, nil)

	if error != nil {
		fmt.Println(error)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("x-api-key", setlistFmApiKey)

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println(error)
	}

	formattedData := utils.FormatJSON(responseBody)

	// clean up memory after execution
	defer response.Body.Close()
	c.String(http.StatusOK, formattedData)

}

func GetArtistsEventsByYear(c *gin.Context) {
	fmt.Println("Making GET request...")
	fmt.Println("Searching events...")
	setlistFmApiKey, _ := os.LookupEnv("SETLIST_FM_API_KEY")
	artistMbid := c.Query("artistMbid")
	year := c.Query("year") // shortcut for c.Request.URL.Query().Get("lastname")
	p := c.DefaultQuery("p", "1")

	// make GET request to API to get user by ID
	apiUrl := fmt.Sprintf("https://api.setlist.fm/rest/1.0/search/setlists?artistMbid=%s&year=%s&p=%s", artistMbid, year, p)
	request, error := http.NewRequest("GET", apiUrl, nil)

	if error != nil {
		fmt.Println(error)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("x-api-key", setlistFmApiKey)

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println(error)
	}

	formattedData := utils.FormatJSON(responseBody)

	// clean up memory after execution
	defer response.Body.Close()
	c.String(http.StatusOK, formattedData)
}

func GetEventById(c *gin.Context) {
	fmt.Println("Making GET request...")
	fmt.Println("Searching events...")
	setlistFmApiKey, _ := os.LookupEnv("SETLIST_FM_API_KEY")
	event_id := c.Param("event_id")

	// make GET request to API to get user by ID
	apiUrl := fmt.Sprintf("https://api.setlist.fm/rest/1.0/setlist/%s", event_id)
	request, error := http.NewRequest("GET", apiUrl, nil)

	if error != nil {
		fmt.Println(error)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("x-api-key", setlistFmApiKey)

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println(error)
	}

	formattedData := utils.FormatJSON(responseBody)

	// clean up memory after execution
	defer response.Body.Close()
	c.String(http.StatusOK, formattedData)

}
