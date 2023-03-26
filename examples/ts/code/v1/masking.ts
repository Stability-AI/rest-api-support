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
const formData = new FormData()
formData.append('init_image', fs.readFileSync('../init_image.png'))
formData.append('mask_image', fs.readFileSync('../mask_image_white.png'))
formData.append('mask_source', 'MASK_IMAGE_WHITE')
formData.append(
  'text_prompts[0][text]',
  'A large spiral galaxy with a bright central bulge and a ring of stars around it'
)
formData.append('cfg_scale', '7')
formData.append('clip_guidance_preset', 'FAST_BLUE')
formData.append('samples', 1)
formData.append('steps', 30)

const response = await fetch(
  `${apiHost}/v1/generation/${engineId}/image-to-image/masking`,
  {
    method: 'POST',
    headers: {
      ...formData.getHeaders(),
      Accept: 'application/json',
      Authorization: `Bearer ${apiKey}`,
    },
    body: formData,
  }
)

if (!response.ok) {
  throw new Error(`Non-200 response: ${await response.text()}`)
}

interface GenerationResponse {
  artifacts: Array<{
    base64: string
    seed: number
    finishReason: string
  }>
}

const responseJSON = (await response.json()) as GenerationResponse

responseJSON.artifacts.forEach((image, index) => {
  fs.writeFileSync(
    `out/v1_img2img_masking_${index}.png`,
    Buffer.from(image.base64, 'base64')
  )
})
