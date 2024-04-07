import { useEffect, useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { API_URL } from './utils/api'

function App() {
  const [count, setCount] = useState(0)
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
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>

      </div>
      <p>
        Responses from the API go below
      </p>
      <pre>
        {JSON.stringify(data, null, 2)}
      </pre>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
