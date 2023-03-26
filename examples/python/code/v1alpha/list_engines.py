import os
import requests

api_host = os.getenv('API_HOST', 'https://api.stability.ai')
url = f"{api_host}/v1alpha/engines/list"

apiKey = os.getenv("STABILITY_API_KEY")
if apiKey is None:
    raise Exception("Missing Stability API key.")

response = requests.get(url, headers={
    "Authorization": apiKey
})

if response.status_code != 200:
    raise Exception("Non-200 response: " + str(response.text))

# Do something with the payload...
payload = response.json()

