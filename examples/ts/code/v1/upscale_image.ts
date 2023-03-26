import fetch from 'node-fetch'
import FormData from 'form-data'
import fs from 'node:fs'

const engineId = 'esrgan-v1-x2plus'
const apiHost = process.env.API_HOST ?? 'https://api.stability.ai'
const apiKey = process.env.STABILITY_API_KEY

if (!apiKey) throw new Error('Missing Stability API key.')

// NOTE: This example is using a NodeJS FormData library.
// Browsers should use their native FormData class.
// React Native apps should also use their native FormData class.
const formData = new FormData()
formData.append('image', fs.readFileSync('../init_image.png'))
formData.append('width', 1024)

const response = await fetch(
  `${apiHost}/v1/generation/${engineId}/image-to-image/upscale`,
  {
    method: 'POST',
    headers: {
      ...formData.getHeaders(),
      Accept: 'image/png',
      Authorization: `Bearer ${apiKey}`,
    },
    body: formData,
  }
)

if (!response.ok) {
  throw new Error(`Non-200 response: ${await response.text()}`)
}

const image = await response.arrayBuffer()
fs.writeFileSync('./out/v1_upscaled_image.png', Buffer.from(image))
