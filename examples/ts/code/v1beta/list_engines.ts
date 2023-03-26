import fetch from 'node-fetch'

const apiHost = process.env.API_HOST ?? 'https://api.stability.ai'
const url = `${apiHost}/v1beta/engines/list`

const apiKey = process.env.STABILITY_API_KEY
if (!apiKey) throw new Error('Missing Stability API key.')

const response = await fetch(url, {
  method: 'GET',
  headers: {
    Authorization: `Bearer ${apiKey}`,
  },
})

if (!response.ok) {
  throw new Error(`Non-200 response: ${await response.text()}`)
}

interface Payload {
  engines: Array<{
    id: string
    name: string
    description: string
    type: string
  }>
}

// Do something with the payload...
const payload = (await response.json()) as Payload
