import json
import os
import requests

engine_id = "stable-inpainting-512-v2-0"
api_host = os.getenv('API_HOST', 'https://api.stability.ai')
url = f"{api_host}/v1alpha/generation/{engine_id}/image-to-image/masking"

init_image = "../init_image.png"
mask_image = "../mask_image_white.png"
output_file = "./out/v1alpha_img2img_masking.png"

api_key = os.getenv("STABILITY_API_KEY")
if api_key is None:
    raise Exception("Missing Stability API key.")

options = json.dumps(
    {
        "mask_source": "MASK_IMAGE_WHITE",
        "cfg_scale": 7,
        "clip_guidance_preset": "FAST_BLUE",
        "height": 512,
        "width": 512,
        "samples": 1,
        "steps": 50,
        "text_prompts": [
            {
                "text": "A large spiral galaxy dog with a bright central bulge and a ring of stars around it",
                "weight": 1
            }
        ],
    }
)

headers = {
    'accept': 'image/png',
    'Authorization': api_key,
}

files = {
    'init_image': open(init_image, 'rb'),
    'mask_image': open(mask_image, 'rb'),
    'options': (None, options),
}

response = requests.post(url, headers=headers, files=files)

if response.status_code != 200:
    raise Exception("Non-200 response: " + str(response.text))

# Write the bytes from response.content to a file
with open(output_file, "wb") as f:
    f.write(response.content)
