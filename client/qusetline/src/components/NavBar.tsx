// import React from "react"
import { Link } from "react-router-dom"
import '../App.css'

function NavBar() {
  return(
  <nav className="nav-bar">
    <Link to='/games_catalog' title="Catalog" className="nav-links">Catalog</Link>
    <Link to='/signup' title="Sign Up" className="nav-links">Sign Up!</Link>
  </nav>
  )
}

export default NavBar