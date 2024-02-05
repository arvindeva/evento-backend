package handlers

import (
	"evento/search-api/utils"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetArtistsByTerm(c *gin.Context) {
	fmt.Println("Making GET request...")
	fmt.Println("Searching artists...")
	term := c.Param("term")
	setlistFmApiKey, _ := os.LookupEnv("SETLIST_FM_API_KEY")

	// make GET request to API to get user by ID
	apiUrl := fmt.Sprintf("https://api.setlist.fm/rest/1.0/search/artists?artistName=%s&sort=relevance", term)
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
