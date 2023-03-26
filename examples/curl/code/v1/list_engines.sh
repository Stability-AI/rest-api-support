if [ -z "$STABILITY_API_KEY" ]; then
    echo "STABILITY_API_KEY environment variable is not set"
    exit 1
fi

BASE_URL=${API_HOST:-https://api.stability.ai}
URL="$BASE_URL/v1/engines/list"

curl -f -sS "$URL" \
  -H 'Accept: application/json' \
  -H "Authorization: Bearer $STABILITY_API_KEY"
