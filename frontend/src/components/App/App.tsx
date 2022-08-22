import React from 'react';
import { BrowserRouter, Route, Routes, useLocation, Navigate } from 'react-router-dom';
import Layout from '../Layout/Layout';
import Dashboard from '../Dashboard/Dashboard';
import Login from '../Login/Login'
import Credentials from '../Auth/authInterfaces';
import Preferences from '../Preferences/Preferences';
import Accounts from '../Accounts/Accounts.js';
import Friends from '../Friends/Friends.js';
import Transactions from '../Transactions/Transactions';
import Register from '../Register/Register';
import Nav from '../Nav/Nav';
import ErrorBoundary from '../ErrorBoundary/ErrorBoundary';
import { useAuth, AuthProvider } from '../Auth/Auth';

function App() {
  return (
    <Layout>
      <AuthProvider>
        <BrowserRouter>
          <ConditionalNav></ConditionalNav>
          <ErrorBoundary>
            <Routes>
              <Route path="/" element={<h1>Placeholder homepage</h1>} />
              <Route path="*" element={<Login />} />
              <Route path="/login" element={<Login />} />
              <Route path="/register" element={<Register />} />
              <Route path="/dashboard" element={<RequireAuth><Dashboard /></RequireAuth>} />
              <Route path="/preferences" element={<RequireAuth><Preferences /></RequireAuth>} />
              <Route path="/accounts" element={<RequireAuth><Accounts /></RequireAuth>} />
              <Route path="/friends" element={<RequireAuth><Friends /></RequireAuth>} />
              <Route path="/transactions" element={<RequireAuth><Transactions /></RequireAuth>} />
            </Routes>
          </ErrorBoundary>
        </BrowserRouter>
      </AuthProvider>
    </Layout>
  );
}

// @ts-ignore
function RequireAuth({ children }) {
  let auth = useAuth();
  let location = useLocation();

  // @ts-ignore
  if (auth.email == '') {
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  return children
}

function ConditionalNav() {
  let auth = useAuth();
  // @ts-ignore
  return auth.email && <Nav></Nav>
}

export default App;
