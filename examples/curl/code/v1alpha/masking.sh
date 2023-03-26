#!/bin/sh

set -e

if [ -z "$STABILITY_API_KEY" ]; then
    echo "STABILITY_API_KEY environment variable is not set"
    exit 1
fi

OUTPUT_FILE=./out/v1alpha_img2img_masking.png
BASE_URL=${API_HOST:-https://api.stability.ai}
URL="$BASE_URL/v1alpha/generation/stable-inpainting-512-v2-0/image-to-image/masking"

curl -f -sS -X POST "$URL" \
  -H 'Content-Type: multipart/form-data' \
  -H 'Accept: image/png' \
  -H "Authorization: $STABILITY_API_KEY" \
  -F 'init_image=@"../init_image.png"' \
  -F 'mask_image=@"../mask_image_white.png"' \
  -F 'options="{
    \"mask_source\": \"MASK_IMAGE_WHITE\",
    \"cfg_scale\": 7,
    \"clip_guidance_preset\": \"FAST_BLUE\",
    \"height\": 512,
    \"width\": 512,
    \"samples\": 1,
    \"steps\": 50,
    \"text_prompts\": [{
        \"text\": \"A large spiral galaxy dog with a bright central bulge and a ring of stars around it\",
        \"weight\": 1
      }]
    }"' \
  -o "$OUTPUT_FILE"
