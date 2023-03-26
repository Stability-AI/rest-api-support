if [ -z "$STABILITY_API_KEY" ]; then
    echo "STABILITY_API_KEY environment variable is not set"
    exit 1
fi

OUTPUT_FILE=./out/v1alpha_txt2img.png
BASE_URL=${API_HOST:-https://api.stability.ai}
URL="$BASE_URL/v1alpha/generation/stable-diffusion-v1-5/text-to-image"

curl -f -sS -X POST "$URL" \
  -H 'Content-Type: application/json' \
  -H 'Accept: image/png' \
  -H "Authorization: $STABILITY_API_KEY" \
  --data-raw '{
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
  }' \
  -o "$OUTPUT_FILE"
