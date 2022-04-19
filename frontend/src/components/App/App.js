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
import { useAuth, AuthProvider } from '../Auth/Auth.js'

function App() {
  // const { token, setToken } = useToken();
  const auth = useAuth();
  return (
    !auth ?
      // Public
      <AuthProvider>
        <BrowserRouter>
          <ErrorBoundary>
            <Routes>
              <Route path="/" element={<h1>Placeholder homepage</h1>} />
              <Route path="*" element={<Login />} />
              <Route path="/login" element={<Login />} />
              <Route path="/register" element={<Register />} />
              {/* <Route path="/" element={<h1>Placeholder homepage</h1>} />
              <Route path="*" setToken={setToken} element={<Login />} />
              <Route path="/login" setToken={setToken} element={<Login />} />
              <Route path="/register" setToken={setToken} element={<Register />} /> */}
            </Routes>
          </ErrorBoundary>
        </BrowserRouter>
      </AuthProvider>
      :
      // Private
      <AuthProvider>
        <BrowserRouter>
          <Nav />
          <ErrorBoundary>
            <Routes>
              <Route path="/" element={<Dashboard />} />
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

export default App;
