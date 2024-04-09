import './App.css'

import { useEffect, useState } from 'react'
import { API_URL } from './utils/api'

function App() {
  const [data, setData] = useState<string | null>(null);

  useEffect(() => {
    fetch(`${API_URL}/posts/`)
      .then((res) => res.text())
      .then((data) => setData(data))
      .catch((err) => {
        console.error(err)
        setData("message: error fetching data. see logs")
      })
  }, [])

  return (
    <main>
      <img src='./favicon.svg' alt="coffee cup" className="logo" />
      <h1>coffee.mykal.codes</h1>
      <p className="read-the-docs">
        Responses from the API go below
      </p>
      <pre>
        {JSON.stringify(data, null, 2)}
      </pre>
    </main>
  )
}

export default App
