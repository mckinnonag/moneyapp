import React, { useState } from 'react';
import { useNavigate, useLocation } from 'react-router-dom'
import useAuth from '../Auth/Auth.js';
import './Login.css'

export default function Login({ setToken }) {
  const [username, setUserName] = useState();
  const [password, setPassword] = useState();
  const { login } = useAuth();
  const location = useLocation();
  const navigate = useNavigate();
  
  const [credentialError, setCredentialError] = useState();
  const from = location.state?.from?.pathname || "/";

  const handleSubmit = async e => {
    e.preventDefault();
    if (username === undefined || password === undefined) {
      setCredentialError('Please enter your username or password.');
      return
    }

    login({username, password}, () => {
      console.log('navigating to ' + from);
      navigate(from, { replace: true });
    });
  }

  return(
      <div className="login-wrapper">
        <h1>Please Log In</h1>
        <form onSubmit={handleSubmit}>
            <label>
                <p>Username</p>
                <input type="email" onChange={e => setUserName(e.target.value)} />
            </label>
            <label>
                <p>Password</p>
                <input type="password" onChange={e => setPassword(e.target.value)} />
            </label>
            <div>
                <button type="submit">Submit</button>
            </div>
        </form>
        <div>
            { credentialError !== '' && <p>{ credentialError }</p> }
        </div>
      </div>
  )
}