import os
import requests

engine_id = "stable-diffusion-v1-5"
api_host = os.getenv('API_HOST', 'https://api.stability.ai')
url = f"{api_host}/v1alpha/generation/{engine_id}/text-to-image"
output_file = "./out/v1alpha_txt2img.png"

apiKey = os.getenv("STABILITY_API_KEY")
if apiKey is None:
    raise Exception("Missing Stability API key.")

payload = {
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
    ],
}

headers = {
    "Content-Type": "application/json",
    "Accept": "image/png",
    "Authorization": apiKey
}

response = requests.post(url, json=payload, headers=headers)

if response.status_code != 200:
    raise Exception("Non-200 response: " + str(response.text))

# Write the bytes from response.content to a file
with open(output_file, "wb") as f:
    f.write(response.content)
