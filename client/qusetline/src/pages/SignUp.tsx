// import React from "react"
import { Link } from "react-router-dom"
// import { useEffect, useState } from "react"

function SignUp() {

  return(
  <div>
    <div className="signup-login-div">
      <br />
      <h1 className="signup-login-header">Sign Up</h1>
      <h5 className="signup-login-link">Already have an account? <Link to='/login' title="Log In!" className="">Log In</Link></h5>
      <form>
        <label htmlFor="username">Username:</label>
        <br />
        <input type="text" id="username" name="username"/>
        <br />
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

export default SignUp