// import { useState } from 'react'
import { Routes, Route } from 'react-router-dom'
import './App.css'
import Catalog from './pages/Catalog.tsx'
import NavBar from './components/NavBar.tsx'

function App() {

  return (
  <>
    <NavBar/>
    <Routes>
      <Route path='/games_catalog' element={<Catalog/>}/>
    </Routes>
  </>
  )
}

export default App
