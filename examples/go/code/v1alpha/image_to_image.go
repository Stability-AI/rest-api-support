package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func main() {
	engineId := "stable-diffusion-v1-5"

	// Build REST endpoint URL
	apiHost, hasApiHost := os.LookupEnv("API_HOST")
	if !hasApiHost {
		apiHost = "https://api.stability.ai"
	}
	reqUrl := apiHost + "/v1alpha/generation/" + engineId + "/image-to-image"

	// Acquire an API key from the environment
	apiKey, hasAPIKey := os.LookupEnv("STABILITY_API_KEY")
	if !hasAPIKey {
		panic("Missing STABILITY_API_KEY environment variable")
	}

	data := &bytes.Buffer{}
	writer := multipart.NewWriter(data)

	// Write the init image to the request
	initImageWriter, _ := writer.CreateFormField("init_image")
	initImageFile, initImageErr := os.Open("../init_image.png")
	if initImageErr != nil {
		panic("Could not open init_image.png")
	}
	_, _ = io.Copy(initImageWriter, initImageFile)

	// Write the options to the request
	optionsFw, _ := writer.CreateFormField("options")
	_, _ = io.Copy(optionsFw, strings.NewReader(`{
		"cfg_scale": 7,
		"clip_guidance_preset": "FAST_BLUE",
		"step_schedule_start": 0.6,
		"step_schedule_end": 0.01,
		"height": 512,
		"width": 512,
		"samples": 1,
		"steps": 50,
		"text_prompts": [{
		  "text": "A large spiral galaxy dog with a bright central bulge and a ring of stars around it",
		  "weight": 1
		}]
  	}`))
	writer.Close()

	// Execute the request
	payload := bytes.NewReader(data.Bytes())
	req, _ := http.NewRequest("POST", reqUrl, payload)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Accept", "image/png")
	req.Header.Add("Authorization", apiKey)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != 200 {
		panic("Non-200 response: " + string(body))
	}

	// Write the bytes to a file
	file, _ := os.Create("./out/v1alpha_img2img.png")
	defer file.Close()
	_, err := file.Write(body)
	if err != nil {
		panic(err)
	}
}
