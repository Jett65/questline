// import { useState } from 'react'
import { Routes, Route } from 'react-router-dom'
import './App.css'
import Catalog from './Pages/Catalog.tsx'
import NavBar from './Components/NavBar.tsx'

function App() {

  return (
  <>
    <NavBar/>
    <Routes>
      <Route path='/games-catalog' element={<Catalog/>}/>
    </Routes>
  </>
  )
}

export default App
