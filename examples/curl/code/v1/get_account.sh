if [ -z "$STABILITY_API_KEY" ]; then
    echo "STABILITY_API_KEY environment variable is not set"
    exit 1
fi

# Determine the URL to use for the request
BASE_URL=${API_HOST:-https://api.stability.ai}
URL="$BASE_URL/v1/user/account"

curl -f -sS "$URL" \
  -H 'Accept: application/json' \
  -H "Authorization: Bearer $STABILITY_API_KEY"
