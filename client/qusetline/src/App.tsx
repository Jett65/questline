// import { useState } from 'react'
import { Routes, Route } from 'react-router-dom'
import './App.css'
import Catalog from './pages/Catalog.tsx'
import NavBar from './components/NavBar.tsx'
import SignUp from './pages/SignUp.tsx'
import LogIn from './pages/LogIn.tsx'
import Home from './pages/Home.tsx'
import GameModes from './pages/Game Modes.tsx'


function App() {

  return (
  <>
    <NavBar/>
    <Routes>
      <Route path='/' element={<Home/>}/>
      <Route path='/games_catalog' element={<Catalog/>}/>
      <Route path='/modes' element={<GameModes/>}/>
      <Route path='/signup' element={<SignUp/>}/>
      <Route path='/login' element={<LogIn/>}/>
    </Routes>
  </>
  )
}

export default App
