package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func FetchCard(queryParams map[string]string, additionalRoute string) (*[]ScryfallCard, error) {
	baseUrl := "https://api.scryfall.com/cards"

	if additionalRoute != "" {
		baseUrl = strings.Join([]string{baseUrl, additionalRoute}, "/")
	}

	var apiUrl string

	if queryParams != nil && len(queryParams) > 0 {
		queryString := buildQueryString(queryParams)
		apiUrl = fmt.Sprintf("%s?q=%s", baseUrl, queryString)
	} else {
		apiUrl = baseUrl
	}

	fmt.Println("Fetching URL:", apiUrl)

	// Perform the GET request
	response, err := http.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cards: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch cards: %s", response.Status)
	}

	// Read the raw response body
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Check if it's a list (contains "data") or a single card
	var rawResponse map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &rawResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// Check if it's a single card (no "data" field)
	if _, found := rawResponse["data"]; !found {
		fmt.Println("Response appears to be a single card")

		var singleCard ScryfallCard
		if err := json.Unmarshal(bodyBytes, &singleCard); err != nil {
			return nil, fmt.Errorf("failed to decode single card: %v", err)
		}

		// Wrap the single card in a slice and return
		cards := []ScryfallCard{singleCard}
		return &cards, nil
	}

	// Handle normal flow with "data" field for multiple cards
	var result struct {
		Data []ScryfallCard `json:"data"`
	}

	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to decode result into ScryfallCard: %v", err)
	}

	if len(result.Data) == 0 {
		return nil, errors.New("no card found")
	}

	return &result.Data, nil
}

func buildQueryString(params map[string]string) string {
	var queryParts []string

	for key, value := range params {
		queryParts = append(queryParts, fmt.Sprintf("%s:%s", url.QueryEscape(key), url.QueryEscape(value)))
	}

	return strings.Join(queryParts, "&")
}
