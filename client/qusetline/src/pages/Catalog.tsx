import React from "react"
import { useEffect, useState } from "react"

function Catalog() {

  const [games, setGames] = useState<Array<{
    id: string,
    name: string,
    image_url: string,
    description: string
  }>>([])

  useEffect(() => {
    fetchGames()
  }, [])
  
  async function fetchGames() {
    let api = 'http://localhost:3006/api/v1'
    
    try {
      const response = await fetch(`${api}/catalog_game`)
      const result = await response.json()
      console.log(result)
      setGames(result)
    } catch (error) {
      console.error(error)
    }
  }
  console.log(games)


  return(
    
  <div>
    <h1>Games Catalog!</h1>
    <h3>Check out our Games</h3>

    <div className="catalogGameDiv">
      {games.map((game) => (
        <div key={game.id} className="gameDivs">{game.name}</div>
      ))}
    </div>
  </div>
  )
}

export default Catalog