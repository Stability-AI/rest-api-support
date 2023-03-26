package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	engineId := "esrgan-v1-x2plus"

	// Build REST endpoint URL
	apiHost, hasApiHost := os.LookupEnv("API_HOST")
	if !hasApiHost {
		apiHost = "https://api.stability.ai"
	}
	reqUrl := apiHost + "/v1/generation/" + engineId + "/image-to-image/upscale"

	// Acquire an API key from the environment
	apiKey, hasAPIKey := os.LookupEnv("STABILITY_API_KEY")
	if !hasAPIKey {
		panic("Missing STABILITY_API_KEY environment variable")
	}

	data := &bytes.Buffer{}
	writer := multipart.NewWriter(data)

	// Write the init image to the request
	initImageWriter, _ := writer.CreateFormField("image")
	initImageFile, initImageErr := os.Open("../init_image.png")
	if initImageErr != nil {
		panic("Could not open init_image.png")
	}
	_, _ = io.Copy(initImageWriter, initImageFile)

	// Write the options to the request
	_ = writer.WriteField("width", "1024")
	writer.Close()

	// Execute the request
	payload := bytes.NewReader(data.Bytes())
	req, _ := http.NewRequest("POST", reqUrl, payload)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Accept", "image/png")
	req.Header.Add("Authorization", "Bearer "+apiKey)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		var body map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			panic(err)
		}
		panic(fmt.Sprintf("Non-200 response: %s", body))
	}

	// Write the response to a file
	out, err := os.Create("./out/v1_upscaled_image.png")
	defer out.Close()
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(out, res.Body)
	if err != nil {
		panic(err)
	}
}
