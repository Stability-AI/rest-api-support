import base64
import os
import requests

engine_id = "stable-inpainting-512-v2-0"
api_host = os.getenv('API_HOST', 'https://api.stability.ai')
api_key = os.getenv("STABILITY_API_KEY")

if api_key is None:
    raise Exception("Missing Stability API key.")

response = requests.post(
    f"{api_host}/v1beta/generation/{engine_id}/image-to-image/masking",
    headers={
        "Accept": 'application/json',
        "Authorization": f"Bearer {api_key}"
    },
    files={
        'init_image': open("../init_image.png", 'rb'),
        'mask_image': open("../mask_image_white.png", 'rb'),
    },
    data={
        "mask_source": "MASK_IMAGE_WHITE",
        "text_prompts[0][text]": "A large spiral galaxy with a bright central bulge and a ring of stars around it",
        "cfg_scale": 7,
        "clip_guidance_preset": "FAST_BLUE",
        "height": 512,
        "width": 512,
        "samples": 1,
        "steps": 50,
    }
)

if response.status_code != 200:
    raise Exception("Non-200 response: " + str(response.text))

data = response.json()

for i, image in enumerate(data["artifacts"]):
    with open(f"./out/v1beta_img2img_masking_{i}.png", "wb") as f:
        f.write(base64.b64decode(image["base64"]))
