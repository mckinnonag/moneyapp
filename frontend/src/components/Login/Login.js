import React, { useState } from 'react';
import PropTypes from 'prop-types';
import './Login.css'

async function loginUser(credentials) {
    return fetch('http://localhost:8080/api/public/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(credentials)
    })
      .then(data => data.json())
}

export default function Login({ setToken }) {
  const [username, setUserName] = useState();
  const [password, setPassword] = useState();
  const [credentialError, setCredentialError] = useState();

  const handleSubmit = async e => {
    e.preventDefault();
    if (username === undefined || password === undefined) {
      setCredentialError('Please enter your username or password.');
      return
    }

    try {
      const token = await loginUser({
        username,
        password
      });
      setToken(token);
    } catch {
      setCredentialError('Incorrect username or password.');
    }
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

Login.propTypes = {
    setToken: PropTypes.func.isRequired
}