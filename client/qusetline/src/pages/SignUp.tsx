// import React from "react"
import { Link, useNavigate } from "react-router-dom"
import React, {  useState } from "react"

function SignUp({setToken}:{setToken:any}) {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const navigate = useNavigate()

  async function handleSubmit (event:React.FormEvent) {
    event.preventDefault();
    let api = 'http://localhost:3006/api/v1'

    try{
      const response = await fetch(`${api}/create_account`, {
        method:'POST',
        headers:{
          Accept: 'application/json',
          'Content-Type':'application/json',
        },
        credentials: "same-origin",
        body: JSON.stringify({
          username,
          email,
          password
        })
      });
      console.log(response.headers)
      const result = await response.json();
      console.log(result.headers);
      setToken(result.token)
      navigate('/');
    }
    catch(err){
      console.error(err);
    }
  }

  return(
  <div>
    <div className="signup-login-div">
      <br />
      <h1 className="signup-login-header">Sign Up</h1>
      <h5 className="signup-login-link">Already have an account? <Link to='/login' title="Log In!" className="">Log In</Link></h5>
      <form onSubmit={handleSubmit}>
        <label htmlFor="username">Username:</label>
        <br />
        <input type="text" id="username" name="username" placeholder="Username" required value={username} onChange={(e) => setUsername(e.target.value)}/>
        <br />
        <label htmlFor="email">Email:</label>
        <br />
        <input type="email" id="email" name="email" placeholder="Email Address" required value={email} onChange={(e) => setEmail(e.target.value)}/>
        <br />
        <label htmlFor="password">Password:</label>
        <br />
        <input type="password" id="password" name="password" placeholder="Password" required value={password} onChange={(e) => setPassword(e.target.value)}/>
        <br />
        <br />
        <button type="submit" className="signup-login-button">SIGN UP</button>
        <br />
        <br />
      </form>
    </div>
  </div>
  )
}

export default SignUp