if [ -z "$STABILITY_API_KEY" ]; then
    echo "STABILITY_API_KEY environment variable is not set"
    exit 1
fi

OUTPUT_FILE=./out/v1_upscaled_image.png
BASE_URL=${API_HOST:-https://api.stability.ai}
URL="$BASE_URL/v1/generation/esrgan-v1-x2plus/image-to-image/upscale"

curl -f -sS -X POST "$URL" \
  -H 'Content-Type: multipart/form-data' \
  -H 'Accept: image/png' \
  -H "Authorization: Bearer $STABILITY_API_KEY" \
  -F 'image=@"../init_image.png"' \
  -F 'width=1024' \
  -o "$OUTPUT_FILE"
