// import React from "react"
import { Link } from "react-router-dom"
import '../App.css'

function NavBar() {
  return(
  <nav className="nav-bar">
    <Link to='/' title="Home" className="nav-links">Home</Link>
    <Link to='/games_catalog' title="Catalog" className="nav-links">Catalog</Link>
    <Link to='/modes' title="Modes" className="nav-links">Game Modes</Link>
    <Link to='/signup' title="Sign Up" className="nav-links">Sign Up!</Link>
  </nav>
  )
}

export default NavBar