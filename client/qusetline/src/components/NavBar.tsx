import { Link } from "react-router-dom"
import '../App.css'

function NavBar({token}:{token:any}) {
  return(
  <nav className="nav-bar">
    <Link to='/' title="Home" className="nav-links">Home</Link>
    <Link to='/games_catalog' title="Catalog" className="nav-links">Catalog</Link>
    <Link to='/modes' title="Modes" className="nav-links">Game Modes</Link>
    {token ?
      <Link to='/myaccount' title="My Account" className="nav-links">Account</Link>
      :
      <Link to='/signup' title="Sign Up" className="nav-links">Sign Up!</Link>
    }

  </nav>
  )
}

export default NavBar