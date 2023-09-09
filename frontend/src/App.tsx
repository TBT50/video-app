import { useState, useEffect } from "react";

import "./App.css";

function App() {
  const [num, setNum] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState("");

  useEffect(() => {
    setLoading(true);
    fetch(
      `https://www.random.org/integers/?num=1&min=1&max=100&col=1&base=10&format=plain&rnd=new`
    )
      .then((response) => {
        if (response.status !== 200) {
          throw new Error(`Something went wrong. Try again.`);
        }

        return response.text();
      })
      .then((random) => {
        setLoading(false);
        setNum(parseInt(random, 10));
      })
      .catch((error) => setError(error.message));
  }, []);

  if (error) return <p>{error}</p>;

  return (
    <>
      <div>
        <p>{loading ? "LOADING..." : num}</p>
      </div>
    </>
  );
}

export default App;
