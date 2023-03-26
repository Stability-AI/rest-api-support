package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	// Build REST endpoint URL
	apiHost, hasApiHost := os.LookupEnv("API_HOST")
	if !hasApiHost {
		apiHost = "https://api.stability.ai"
	}
	reqUrl := apiHost + "/v1alpha/user/balance"

	// Acquire an API key from the environment
	apiKey, hasAPIKey := os.LookupEnv("STABILITY_API_KEY")
	if !hasAPIKey {
		panic("Missing STABILITY_API_KEY environment variable")
	}

	// Build the request
	req, _ := http.NewRequest("GET", reqUrl, nil)
	req.Header.Add("Authorization", apiKey)

	// Execute the request
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != 200 {
		panic("Non-200 response: " + string(body))
	}

	// Do something with the payload...
	// payload := string(body)
}
