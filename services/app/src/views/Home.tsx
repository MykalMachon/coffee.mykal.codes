import { useEffect, useState } from "react";
import { API_URL } from "../utils/api";
import Layout from "../components/Layout";

const HomePage = () => {
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
    <Layout>
      <header className="header--primary container">
        <h1>Mykal's coffee</h1>
        <p>A site for mykal to talk about his silly little coffee obsession and learn tech at the same time.</p>
        <p>~</p>
      </header>
      <main className="container">
        <p>some stuff coming soon. example data below</p>
        <pre>
          {JSON.stringify(data, null, 2)}
        </pre>
      </main>
    </Layout>
  );
}

export default HomePage;