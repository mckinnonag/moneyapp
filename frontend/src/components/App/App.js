import React from 'react';
import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Dashboard from '../Dashboard/Dashboard';
import Login from '../Login/Login'
import Preferences from '../Preferences/Preferences';
import Accounts from '../Accounts/Accounts.js'
import Friends from '../Friends/Friends.js'
import useToken from './useToken';
import ResponsiveAppBar from '../ResponsiveAppBar/ResponsiveAppBar.js'

function App() {
  const { token, setToken } = useToken();

  if (!token) {
    return <Login setToken={setToken} />
  }

  return (
    <div className="wrapper">
      <BrowserRouter>
        <ResponsiveAppBar />
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/preferences" element={<Preferences />} />
          <Route path="/accounts" element={<Accounts />} />
          <Route path="/friends" element={<Friends />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
