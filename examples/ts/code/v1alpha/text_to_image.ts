import fetch from 'node-fetch'
import fs from 'node:fs'

const engineId = 'stable-diffusion-v1-5'
const apiHost = process.env.API_HOST ?? 'https://api.stability.ai'
const url = `${apiHost}/v1alpha/generation/${engineId}/text-to-image`
const apiKey = process.env.STABILITY_API_KEY

if (!apiKey) throw new Error('Missing Stability API key.')

const response = await fetch(url, {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
    Accept: 'image/png',
    Authorization: apiKey,
  },
  body: JSON.stringify({
    cfg_scale: 7,
    clip_guidance_preset: 'FAST_BLUE',
    height: 512,
    width: 512,
    samples: 1,
    steps: 50,
    text_prompts: [
      {
        text: 'A lighthouse on a cliff',
        weight: 1,
      },
    ],
  }),
})

if (!response.ok) {
  throw new Error(`Non-200 response: ${await response.text()}`)
}

try {
  const writeStream = fs.createWriteStream(`./out/v1alpha_txt2img.png`)
  response.body?.pipe(writeStream)
} catch (e) {
  console.error(e)
}
