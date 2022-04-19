import React, { useState, useContext, createContext } from 'react';
import useToken from './useToken.js'

const authContext = createContext(null);

export function AuthProvider({ children }) {
    // const { token, setToken } = useToken();
  const auth = useProvideAuth();
  return <authContext.Provider value={auth}>{children}</authContext.Provider>;
}

export const useAuth = () => {
  return useContext(authContext);
}

function useProvideAuth() {
  const [user, setUser] = useState(null);

  const login = (credentials, callback) => {
    return fakeAuthProvider.login(() => {
      setUser(credentials.email);
      callback();
    });
  };

  const logout = (callback) => {
    return fakeAuthProvider.logout(() => {
      setUser(null);
      callback();
    });
  };

  return {
    user,
    login,
    logout,
  };
}

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

const fakeAuthProvider = {
    isAuthenticated: false,
    login(credentials, setToken, callback) {
      const token = loginUser(credentials)
      setToken(token);
      fakeAuthProvider.isAuthenticated = true;
      setTimeout(callback, 100); // fake async
    },
    logout(callback) {
      fakeAuthProvider.isAuthenticated = false;
      setTimeout(callback, 100);
    },
  };

export default function AuthConsumer() {
  return useContext(authContext);
}