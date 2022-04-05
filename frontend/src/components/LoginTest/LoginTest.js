import React, { useState } from 'react';
import { useNavigate } from "react-router-dom";
import useAuth from '../Auth/Auth.js';

const LoginTest = () => {
    const navigate = useNavigate();
    const { login } = useAuth();
  
    const handleLogin = () => {
      login().then(() => {
        navigate("/dashboard");
      });
    };
  
    return (
      <div>
        <h1>Login</h1>
        <button onClick={handleLogin}>Log in</button>
      </div>
    );
  };

export default LoginTest;