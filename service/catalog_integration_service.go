package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Product represents the structure of each product in the response data
type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Qty         int    `json:"qty"`
	Active      bool   `json:"active"`
	Deleted     bool   `json:"deleted"`
}

// Response represents the entire structure of the API response
type Response struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	Signature *string   `json:"signature"`
	Data      []Product `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

// FetchProducts sends a GET request to the given URL and fetches the products data
func FetchProducts(baseURL string, codes []string) ([]Product, error) {
	// Build the query string
	query := url.Values{}
	query.Set("codes", strings.Join(codes, ","))

	// Construct the full URL
	fullURL := fmt.Sprintf("%s?%s", baseURL, query.Encode())
	log.Printf("Request URL: %s", fullURL)

	// Send the GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	// Decode the response body
	var apiResponse Response
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Return the product data
	return apiResponse.Data, nil
}
