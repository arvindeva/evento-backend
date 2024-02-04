package handlers

import (
	"evento/search-api/utils"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetEventsByArtistId(mbid string, p string) string {
	fmt.Println("Making GET request...")
	fmt.Println("Searching events...")
	setlistFmApiKey, _ := os.LookupEnv("SETLIST_FM_API_KEY")

	// make GET request to API to get user by ID
	apiUrl := fmt.Sprintf("https://api.setlist.fm/rest/1.0/artist/%s/setlists?p=%s", mbid, p)
	request, error := http.NewRequest("GET", apiUrl, nil)

	if error != nil {
		fmt.Println(error)
		return error.Error()
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
	return formattedData
}

func GetArtistsEventsByYear(artistMbid string, year string, p string) string {
	fmt.Println("Making GET request...")
	fmt.Println("Searching events...")
	setlistFmApiKey, _ := os.LookupEnv("SETLIST_FM_API_KEY")

	// make GET request to API to get user by ID
	apiUrl := fmt.Sprintf("https://api.setlist.fm/rest/1.0/search/setlists?artistMbid=%s&year=%s&p=%s", artistMbid, year, p)
	request, error := http.NewRequest("GET", apiUrl, nil)

	if error != nil {
		fmt.Println(error)
		return error.Error()
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
	return formattedData
}
