import React from "react"
import { useEffect, useState } from "react"

function Catalog() {
  const [games, setGames] = useState([]);

  useEffect(() => {
    // fetchGames()
  }, [])

  async function fetchGames () {
    let api = 'http://localhost:3006/api/v1'

    const response = await fetch(`${api}/catalog_game`)
    const result = await response.json()

    console.log(result)
    setGames(result)
  }

  return(
  <div>
    <h1>Games Catalog!</h1>
    <h3>Check out our Games</h3>

    <div className="catalogGameDiv">
      {/* {games.map((game) => (
        <div key={game.id}></div>
      ))} */}
      <button onClick={fetchGames}>Click me For games</button>
    </div>
  </div>
  )
}

export default Catalog