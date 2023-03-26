import os
import requests

engine_id = "esrgan-v1-x2plus"
api_host = os.getenv("API_HOST", "https://api.stability.ai")
api_key = os.getenv("STABILITY_API_KEY")

if api_key is None:
    raise Exception("Missing Stability API key.")

response = requests.post(
    f"{api_host}/v1/generation/{engine_id}/image-to-image/upscale",
    headers={
        "Accept": "image/png",
        "Authorization": f"Bearer {api_key}"
    },
    files={
        "image": open("../init_image.png", "rb")
    },
    data={
        "width": 1024,
    }
)

if response.status_code != 200:
    raise Exception("Non-200 response: " + str(response.text))

with open(f"./out/v1_upscaled_image.png", "wb") as f:
    f.write(response.content)
