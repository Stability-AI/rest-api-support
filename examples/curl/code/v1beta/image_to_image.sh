if [ -z "$STABILITY_API_KEY" ]; then
    echo "STABILITY_API_KEY environment variable is not set"
    exit 1
fi

OUTPUT_FILE=./out/v1beta_img2img.png
BASE_URL=${API_HOST:-https://api.stability.ai}
URL="$BASE_URL/v1beta/generation/stable-diffusion-v1-5/image-to-image"

curl -f -sS -X POST "$URL" \
  -H 'Content-Type: multipart/form-data' \
  -H 'Accept: image/png' \
  -H "Authorization: Bearer $STABILITY_API_KEY" \
  -F 'init_image=@"../init_image.png"' \
  -F 'init_image_mode=IMAGE_STRENGTH' \
  -F 'image_strength=0.35' \
  -F 'text_prompts[0][text]=A galactic dog in space' \
  -F 'cfg_scale=7' \
  -F 'clip_guidance_preset=FAST_BLUE' \
  -F 'height=512' \
  -F 'width=512' \
  -F 'samples=1' \
  -F 'steps=50' \
  -o "$OUTPUT_FILE"
