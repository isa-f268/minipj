package utils

import (
	"encoding/json"
	"net/http"
)

type Books struct {
	Title   string `json:"title"`
	Authors []struct {
		Name string `json:"name"`
	} `json:"authors"`
}
type OpenLibraryResponse struct {
	Works []Books `json:"works"`
}

func GetInterNationalBooks() (OpenLibraryResponse, error) {
	var result OpenLibraryResponse
	resp, err := http.Get("https://openlibrary.org/subjects/mystery.json?limit=15")

	if err != nil {
		return OpenLibraryResponse{}, err
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return OpenLibraryResponse{}, err
	}
	return result, nil
}
