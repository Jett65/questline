// import { useState } from 'react'
import { Routes, Route, useNavigate } from 'react-router-dom'
import { useState, useEffect} from 'react'
import './App.css'
import Catalog from './pages/Catalog.tsx'
import NavBar from './components/NavBar.tsx'
import SignUp from './pages/SignUp.tsx'
import LogIn from './pages/LogIn.tsx'
import Home from './pages/Home.tsx'
import GameModes from './pages/Game Modes.tsx'
import MyAccount from './pages/MyAccount.tsx'


function App() {
  const [token, setToken] = useState("");
  const navigate = useNavigate();

  function isSignedIn(token:string){
    if (!token){
      if(document.cookie){
        setToken(document.cookie.slice(4))
      }
    }
  };

  useEffect(() => {
    navigate("/");
    isSignedIn(token)
  }, [token])

  return (
  <>
    <NavBar token={token}/>
    <Routes>
      <Route path='/' element={<Home/>}/>
      <Route path='/games_catalog' element={<Catalog/>}/>
      <Route path='/modes' element={<GameModes/>}/>
      {token ?
        <>
          <Route path='/myaccount' element={<MyAccount/>}/>
        </>
      :
        <>
          <Route path='/signup' element={<SignUp setToken={setToken}/>}/>
          <Route path='/login' element={<LogIn setToken={setToken}/>}/>
        </>
      }
    </Routes>
  </>
  )
}

export default App
