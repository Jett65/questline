// import React from "react"
import { Link, } from "react-router-dom"
import React, {  useState } from "react"

function SignUp({setToken}:{setToken:any}) {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [focusedUsername, setFocusedUsername] = useState(false);
  const [focusedEmail, setFocusedEmail] = useState(false);
  const [focusedPassword, setFocusedPassword] = useState(false);
  const errorMessage = [
    "*Username must be 4-16 characters long with no special characters, letters and numbers only*",
    "*Must be valid Email Address*",
    "*Password must be 8-16 characters long with at least 1 letter, number, and special character*"
  ]

  // Possible function for handling focused
  // const handleFocused = (focusName: string) => {
  //   focused.username = true;
  //   console.log(focusName)
  // };


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
        credentials: "include",
        body: JSON.stringify({
          username,
          email,
          password
        })
      });
      const result = await response.json();
      let resultToken = document.cookie.slice(4);
      setToken(resultToken)
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
        <input type="text" id="username" name="username" placeholder="Username" pattern="^[A-Za-z0-9]{4,16}$" className="signup-input" required value={username} data-focused={focusedUsername.toString()} onBlur={() => setFocusedUsername(true)} onChange={(e) => setUsername(e.target.value)}/>
        <br />
        <span id="username">{errorMessage[0]}</span>
        <br />
        <label htmlFor="email">Email:</label>
        <br />
        <input type="email" id="email" name="email" placeholder="Email Address" className="signup-input" required value={email} data-focused={focusedEmail.toString()} onBlur={() => setFocusedEmail(true)} onChange={(e) => setEmail(e.target.value)}/>
        <br />
        <span id="email">{errorMessage[1]}</span>
        <br />
        <label htmlFor="password">Password:</label>
        <br />
        <input type="password" id="password" name="password" placeholder="Password" className="signup-input" pattern="^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$" required value={password} data-focused={focusedPassword.toString()} onBlur={() => setFocusedPassword(true)} onChange={(e) => setPassword(e.target.value)}/>
        <br />
        <span id="password">{errorMessage[2]}</span>
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
