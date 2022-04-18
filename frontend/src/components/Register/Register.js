import React, { useState } from 'react';
import PropTypes from 'prop-types';
import useAuth from '../Auth/Auth.js';

async function registerUser(credentials) {
    return fetch('http://localhost:8080/api/public/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(credentials)
    })
      .then(data => data.json())
}

export default function Register({ setToken }) {
  const [username, setUserName] = useState();
  const [password, setPassword] = useState();
  const [passwordTwo, setPasswordTwo] = useState();
  const [credentialError, setCredentialError] = useState();

  const handleSubmit = async e => {
    e.preventDefault();
    if (username === undefined || password === undefined) {
      setCredentialError('Please enter a username and password.');
      return;
    }
    if (password !== passwordTwo) {
        setCredentialError('Passwords do not match.');
        return;
    }

    try {
      const token = await registerUser({
        username,
        password
      });
      setToken(token);
    } catch (error) {
        setCredentialError(error);
    }
  }

  return(
      <div className="login-wrapper">
        <h1>Please Register</h1>
        <form onSubmit={handleSubmit}>
            <label>
                <div>
                    <input type="email" placeholder="username" onChange={e => setUserName(e.target.value)} />
                </div>
            </label>
            <label>
                <div>
                    <input type="password" placeholder="password" onChange={e => setPassword(e.target.value)} />
                </div>
            </label>
            <label>
                <div>
                    <input type="password" placeholder="confirm password" onChange={e => setPasswordTwo(e.target.value)}/>
                </div>
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

Register.propTypes = {
    setToken: PropTypes.func.isRequired
}