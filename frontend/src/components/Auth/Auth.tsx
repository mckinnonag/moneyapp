import React, { useState, useContext, createContext } from 'react';
import Credentials from './authInterfaces';

const authContext = createContext(null);

// @ts-ignore
export function AuthProvider({ children }) {
  const auth = useProvideAuth();
   // @ts-ignore
  return <authContext.Provider value={auth}>{children}</authContext.Provider>;
}

export const useAuth = () => {
  return useContext(authContext);
}

function useProvideAuth() {
  const [user, setUser] = useState<string | undefined>('');

  const login = (credentials: Credentials, callback: Function) => {
    customAuthProvider.login(credentials, callback);
    setUser(credentials.email);
  };

  const logout = (callback: Function) => {
    return customAuthProvider.logout(() => {
      setUser('');
      callback();
    });
  };

  return {
    user,
    login,
    logout,
  };
}

const customAuthProvider = {
  isAuthenticated: false,
  login: (credentials: Credentials, callback: Function) => {
    const request = new Request('http://localhost:8080/api/public/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(credentials),
    });
    return fetch(request)
      .then(response => {
          if (response.status < 200 || response.status >= 300) {
              throw new Error(response.statusText);
          }
          return response.json();
      })
      .then(token => {
          sessionStorage.setItem('token', JSON.stringify(token));
          customAuthProvider.isAuthenticated = true;
      })
      .then(callback())
      .catch(() => {
          throw new Error('Network error')
      });
  },
  logout(callback: Function) {
    customAuthProvider.isAuthenticated = false;
    sessionStorage.removeItem('token');
  },
};

export default function AuthConsumer() {
  return useContext(authContext);
}