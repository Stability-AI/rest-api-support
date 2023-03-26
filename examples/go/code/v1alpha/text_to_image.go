package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	// Build REST endpoint URL w/ specified engine
	engineId := "stable-diffusion-v1-5"
	apiHost, hasApiHost := os.LookupEnv("API_HOST")
	if !hasApiHost {
		apiHost = "https://api.stability.ai"
	}
	reqUrl := apiHost + "/v1alpha/generation/" + engineId + "/text-to-image"

	// Acquire an API key from the environment
	apiKey, hasAPIKey := os.LookupEnv("STABILITY_API_KEY")
	if !hasAPIKey {
		panic("Missing STABILITY_API_KEY environment variable")
	}

	var data = []byte(`{
		"cfg_scale": 7,
		"clip_guidance_preset": "FAST_BLUE",
		"height": 512,
		"width": 512,
		"samples": 1,
		"steps": 50,
		"text_prompts": [
		  {
			"text": "A lighthouse on a cliff",
			"weight": 1
		  }
		]
  	}`)

	req, _ := http.NewRequest("POST", reqUrl, bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "image/png")
	req.Header.Add("Authorization", apiKey)

	// Execute the request & read all the bytes of the response
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != 200 {
		panic("Non-200 response: " + string(body))
	}

	// Write the bytes to a file
	file, _ := os.Create("./out/v1alpha_txt2img.png")
	defer file.Close()
	_, err := file.Write(body)
	if err != nil {
		panic(err)
	}
}
