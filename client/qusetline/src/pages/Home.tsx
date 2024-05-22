// import React from "react"
import { Link } from "react-router-dom"
// import { useEffect, useState } from "react"

function Home() {

  return(
  <div>
    <h1 className="home-header">Welcome to Qusetline!</h1>
    <hr />
    <div id="home-catalog-div">
      <h4 className="home-catalog-link-header">View our Game Catalog</h4>
      <Link to='/games_catalog' title="catalog" className="home-to-catalog-link">Catalog</Link>
      <br />
    </div>
  </div>
  )
}

export default Home