// import React from "react"
import { Link } from "react-router-dom"
import React, { useState } from "react"


function LogIn(props:any) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

   

  async function handleSubmit (e: React.FormEvent) {
    e.preventDefault();
    let api = 'http://localhost:3006/api/v1'

    try{
      const response = await fetch(`${api}/login`, {
        method:'POST',
        headers:{
          Accept: 'application/json',
          'Content-Type':'application/json',
        },
        credentials: "include",
        body: JSON.stringify({
          username,
          password
        })
      });
      const result = await response.json();
      console.log(result)
      let resultToken = document.cookie.slice(4);
      props.setToken(resultToken)
    } catch(err) {
      console.error(err);
      setError("Username or Password is incorrect")
    }
  }

  return(
  <div>
    <div className="signup-login-div">
      <br />
      <h1 className="signup-login-header">Log In</h1>
      <h5 className="signup-login-link">Don't have an account? <Link to='/signup' title="Log In!" className="">Sign Up</Link></h5>
      {error ?
        <h4 className="incorrect-info-header">{error}</h4>
        :
        <h4></h4>
      }
      <form onSubmit={handleSubmit}>
        <label htmlFor="username">Username:</label>
        <br />
        <input type="username" id="username" name="username" placeholder="Username" required value={username} onChange={(e) => setUsername(e.target.value)}/>
        <br />
        <br />
        <label htmlFor="password">Password:</label>
        <br />
        <input type="password" id="password" name="password" placeholder="Password" required value={password} onChange={(e) => setPassword(e.target.value)}/>
        <br />
        <br />
        <button type="submit" className="signup-login-button">LOG IN</button>
        <br />
        <br />
      </form>
    </div>
  </div>
  )
}

export default LogIn
