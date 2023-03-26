import fetch from 'node-fetch'
import FormData from 'form-data'
import fs from 'node:fs'

const engineId = 'stable-inpainting-512-v2-0'
const apiHost = process.env.API_HOST ?? 'https://api.stability.ai'
const apiKey = process.env.STABILITY_API_KEY

if (!apiKey) throw new Error('Missing Stability API key.')

// NOTE: This example is using a NodeJS FormData library. Browser
// implementations should use their native FormData class. React Native
// implementations should also use their native FormData class.
const form = new FormData()
form.append('init_image', fs.readFileSync('../init_image.png'))
form.append('mask_image', fs.readFileSync('../mask_image_white.png'))
form.append(
  'options',
  JSON.stringify({
    mask_source: 'MASK_IMAGE_WHITE',
    cfg_scale: 7,
    clip_guidance_preset: 'FAST_BLUE',
    height: 512,
    width: 512,
    samples: 1,
    steps: 50,
    text_prompts: [
      {
        text: 'A large spiral galaxy with a bright central bulge and a ring of stars around it',
        weight: 1,
      },
    ],
  })
)

const response = await fetch(
  `${apiHost}/v1alpha/generation/${engineId}/image-to-image/masking`,
  {
    method: 'POST',
    headers: {
      ...form.getHeaders(),
      Accept: 'image/png',
      Authorization: apiKey,
    },
    body: form,
  }
)

if (!response.ok) {
  throw new Error(`Non-200 response: ${await response.text()}`)
}

try {
  const writeStream = fs.createWriteStream(`./out/v1alpha_img2img_masking.png`)
  response.body?.pipe(writeStream)
} catch (e) {
  console.error(e)
}
