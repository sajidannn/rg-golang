package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Quotes struct {
	Tags   []string `json:"tags"`
	Author string   `json:"author"`
	Quote  string   `json:"content"`
}

func ClientGet() ([]Quotes, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", "http://api.quotable.io/quotes/random?limit=3", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var quotes []Quotes
	err = json.Unmarshal(body, &quotes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return quotes, nil
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	requestBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("https://postman-echo.com/post", "application/json", requestBody)
	if err != nil {
		return Postman{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Postman{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var postman Postman
	err = json.Unmarshal(body, &postman)
	if err != nil {
		return Postman{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return postman, nil
}

func main() {
	get, _ := ClientGet()
	fmt.Println(get)

	post, _ := ClientPost()
	fmt.Println(post)
}
