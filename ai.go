package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func prompter(prompt string) (string, error) {
	apiKey, exists := os.LookupEnv("GEMINI_KEY")
	if !exists {
		return "", errors.New("GEMINI_KEY not found. Create a .env file in root and add you key there")
	}
	// Define the request body structure
	requestBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{
						"text": prompt,
					},
				},
			},
		},
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Make the POST request
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent?key=%s", apiKey)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}
	var response Response
	err = json.Unmarshal([]byte(responseBody), &response)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	return string(response.Candidates[0].Content.Parts[0].Text), nil
}
