import React, { useState } from 'react';
import axios from "axios"
function Login() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [loggedIn, setLoggedIn] = useState(false);

  const handleLogin = async (event) => {
    event.preventDefault();
    let data=await axios({
        method:"Post",
        url:"http://localhost:3000/login",
        headers:{},
        data:{
            "username":username,
            "password":password,
        }
    })
    console.log(data)
  };

  return (
    <div>
      {loggedIn ? (
        <div>
          <h2>Welcome, {username}!</h2>
          <button onClick={() => setLoggedIn(false)}>Logout</button>
        </div>
      ) : (
        <form onSubmit={handleLogin}>
          <label>
            Username:
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </label>
          <br />
          <label>
            Password:
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </label>
          <br />
          <button type="submit">Login</button>
        </form>
      )}
    </div>
  );
}

export default Login;
