import React, { useState } from 'react';
import axios from "axios"
function Signup() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const handleLogin =  (event) => {
    event.preventDefault();
     axios({
        method:"Post",
        url:"http://localhost:3000/signup",
        headers:{},
        data:{
            "username":username,
            "password":password,
        }
    }).then((data)=> 
         {if (data.status==202)
             window.location.href="http://localhost:5173/Animelist"
            else{
                alert("User already exists")
            }
    }).catch((err)=>console.log(err))
    
  };

  return (
    <div className=' rounded-lg shadow-lg'>
        <form onSubmit={handleLogin} className=' m-10 flex flex-col gap-3'>
             <div className=' flex gap-4'>
                <span >Username </span>
                <input
                 type="text"
                 value={username}
                 required 
                 onChange={(e)=>setUsername(e.target.value)}/>
            </div>
            <div className=' flex gap-4'>
                <span>Password</span>
                <input 
                type="password"
                value={password}
                required 
                onChange={(e)=>setPassword(e.target.value)}/>
            </div>
            <button type='submit'>Signup</button>
        </form>
    </div>
  );
}

export default Signup;
