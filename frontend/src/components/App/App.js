import React from 'react';
import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Dashboard from '../Dashboard/Dashboard';
import Login from '../Login/Login'
import Preferences from '../Preferences/Preferences';
import Accounts from '../Accounts/Accounts.js'
import Friends from '../Friends/Friends.js'
import Transactions from '../Transactions/Transactions';
import Register from '../Register/Register';
import useToken from '../Auth/useToken';
import Nav from '../Nav/Nav.js'
import ErrorBoundary from '../ErrorBoundary/ErrorBoundary.js'
import { AuthProvider } from '../Auth/Auth.js'

function App() {
  const { token, setToken } = useToken();
  if (!token) {
    return (
      <AuthProvider>
        <BrowserRouter>
          <ErrorBoundary>
            <Login setToken={setToken} />
            <Routes>
              <Route path="/login" setToken={setToken} element={<Login />} />
              <Route path="/register" setToken={setToken} element={<Register />} />
            </Routes>
          </ErrorBoundary>
        </BrowserRouter>
      </AuthProvider>
    )
  } else {
    return (
      <AuthProvider>
        <BrowserRouter>
          <Nav />
          <ErrorBoundary>
            <Routes>
              <Route path="/" element={<Dashboard />} />
              {/* <Route path="/login" setToken={setToken} element={<Login /> } />
              <Route path="/logintest" element={<LoginTest />} />
              <Route path="/register" setToken={setToken} element={<Register />} /> */}
              <Route path="/dashboard" element={<Dashboard />} />
              <Route path="/preferences" element={<Preferences />} />
              <Route path="/accounts" element={<Accounts />} />
              <Route path="/friends" element={<Friends />} />
              <Route path="/transactions" element={<Transactions />} />
            </Routes>
          </ErrorBoundary>
        </BrowserRouter>
      </AuthProvider>
    );
  }
}

export default App;
