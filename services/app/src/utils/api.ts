export const API_URL = import.meta.env.PROD ? '/ap`i' : 'http://localhost:2080'

type FetchMethod = 'GET' | 'POST' | 'PUT' | 'DELETE'
type FetchOptions = {
  method: FetchMethod,
  headers: Record<string, string>
  body?: string
}

export const ApiFetcher = (endpoint: string, options: FetchOptions, token?: string) => {
  return fetch(`${API_URL}/${endpoint}`, {
    ...options,
    headers: {
      ...options.headers,
      Authorization: token ? `Bearer ${token}` : ''
    }
  })
}