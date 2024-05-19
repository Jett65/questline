// import React from "react"
import { Link } from "react-router-dom"
// import { useEffect, useState } from "react"

function LogIn() {

  return(
  <div>
    <div className="signup-login-div">
      <br />
      <h1 className="signup-login-header">Login</h1>
      <h5 className="signup-login-link">Don't have an account? <Link to='/signup' title="Log In!" className="">Sign Up</Link></h5>
      <form>
        <label htmlFor="email">Email:</label>
        <br />
        <input type="text" id="email" name="email"/>
        <br />
        <label htmlFor="password">Password:</label>
        <br />
        <input type="text" id="password" name="password"/>
        <br />
        <br />
      </form>
    </div>
  </div>
  )
}

export default LogIn