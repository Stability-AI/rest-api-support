import fetch from 'node-fetch'
import FormData from 'form-data'
import fs from 'node:fs'

const engineId = 'stable-diffusion-v1-5'
const apiHost = process.env.API_HOST ?? 'https://api.stability.ai'
const apiKey = process.env.STABILITY_API_KEY

if (!apiKey) throw new Error('Missing Stability API key.')

// NOTE: This example is using a NodeJS FormData library. Browser
// implementations should use their native FormData class. React Native
// implementations should also use their native FormData class.
const formData = new FormData()
formData.append('init_image', fs.readFileSync('../init_image.png'))
formData.append(
  'options',
  JSON.stringify({
    cfg_scale: 7,
    clip_guidance_preset: 'FAST_BLUE',
    step_schedule_start: 0.6,
    step_schedule_end: 0.01,
    height: 512,
    width: 512,
    samples: 1,
    steps: 50,
    text_prompts: [
      {
        text: 'Galactic dog',
        weight: 1,
      },
    ],
  })
)

const response = await fetch(
  `${apiHost}/v1alpha/generation/${engineId}/image-to-image`,
  {
    method: 'POST',
    headers: {
      ...formData.getHeaders(),
      Accept: 'image/png',
      Authorization: apiKey,
    },
    body: formData,
  }
)

if (!response.ok) {
  throw new Error(`Non-200 response: ${await response.text()}`)
}

try {
  const writeStream = fs.createWriteStream(`./out/v1alpha_img2img.png`)
  response.body?.pipe(writeStream)
} catch (e) {
  console.error(e)
}
