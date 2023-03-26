import fetch from 'node-fetch'

const apiHost = process.env.API_HOST ?? 'https://api.stability.ai'
const url = `${apiHost}/v1beta/user/account`

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

interface User {
  id: string
  profile_picture: string
  email: string
  organizations?: Array<{
    id: string
    name: string
    role: string
    is_default: boolean
  }>
}

// Do something with the user...
const user = (await response.json()) as User
