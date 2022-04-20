import React, { useState, useContext, createContext } from 'react';

const authContext = createContext(null);

export function AuthProvider({ children }) {
  const auth = useProvideAuth();
  return <authContext.Provider value={auth}>{children}</authContext.Provider>;
}

export const useAuth = () => {
  return useContext(authContext);
}

function useProvideAuth() {
  const [user, setUser] = useState(null);

  const login = (credentials, callback) => {
    customAuthProvider.login(credentials, callback);
    console.log('setting user to ' + credentials.username);
    setUser(credentials.username);
  };

  const logout = (callback) => {
    return customAuthProvider.logout(() => {
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

const customAuthProvider = {
  isAuthenticated: false,
  login: (credentials, callback) => {
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
  logout(callback) {
    customAuthProvider.isAuthenticated = false;
    // setTimeout(callback, 100);
  },
};

export default function AuthConsumer() {
  return useContext(authContext);
}