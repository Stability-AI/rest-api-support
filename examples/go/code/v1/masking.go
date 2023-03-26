package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type MaskingImage struct {
	Base64       string `json:"base64"`
	Seed         uint32 `json:"seed"`
	FinishReason string `json:"finishReason"`
}

type MaskingResponse struct {
	Images []MaskingImage `json:"artifacts"`
}

func main() {
	engineId := "stable-inpainting-512-v2-0"

	// Build REST endpoint URL
	apiHost, hasApiHost := os.LookupEnv("API_HOST")
	if !hasApiHost {
		apiHost = "https://api.stability.ai"
	}
	reqUrl := apiHost + "/v1/generation/" + engineId + "/image-to-image/masking"

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

	// Write the mask image to the request
	maskImageWriter, _ := writer.CreateFormField("mask_image")
	maskImageFile, maskImageErr := os.Open("../mask_image_white.png")
	if maskImageErr != nil {
		panic("Could not open mask_image_white.png")
	}
	_, _ = io.Copy(maskImageWriter, maskImageFile)

	// Write the options to the request
	_ = writer.WriteField("mask_source", "MASK_IMAGE_WHITE")
	_ = writer.WriteField("text_prompts[0][text]", "A large spiral galaxy with a bright central bulge and a ring of stars around it")
	_ = writer.WriteField("cfg_scale", "7")
	_ = writer.WriteField("clip_guidance_preset", "FAST_BLUE")
	_ = writer.WriteField("samples", "1")
	_ = writer.WriteField("steps", "30")
	writer.Close()

	// Execute the request & read all the bytes of the response
	payload := bytes.NewReader(data.Bytes())
	req, _ := http.NewRequest("POST", reqUrl, payload)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Accept", "application/json")
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

	// Decode the JSON body
	var body MaskingResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		panic(err)
	}

	// Write the images to disk
	for i, image := range body.Images {
		outFile := fmt.Sprintf("./out/v1_img2img_masking_%d.png", i)
		file, err := os.Create(outFile)
		if err != nil {
			panic(err)
		}

		imageBytes, err := base64.StdEncoding.DecodeString(image.Base64)
		if err != nil {
			panic(err)
		}

		if _, err := file.Write(imageBytes); err != nil {
			panic(err)
		}

		if err := file.Close(); err != nil {
			panic(err)
		}
	}
}
