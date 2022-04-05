import * as React from 'react';
import useToken from './useToken.js'

const authContext = React.createContext(null);

export function AuthProvider({ children }) {
    const { token, setToken } = useToken('');
//   let [user, setUser] = React.useState(null);

  let login = (credentials, setToken, callback) => {
    return fakeAuthProvider.login(() => {
    //   setUser(newUser);
      callback();
    });
  };

  let logout = (callback) => {
    return fakeAuthProvider.logout(() => {
    //   setUser(null);
      callback();
    });
  };

  let value = { token, login, logout };

  return <authContext.Provider value={value}>{children}</authContext.Provider>;
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
    logout(credentials, callback) {
      fakeAuthProvider.isAuthenticated = false;
      setTimeout(callback, 100);
    },
  };

export default function AuthConsumer() {
  return React.useContext(authContext);
}