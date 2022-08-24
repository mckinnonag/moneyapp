import React, { useState, useContext } from 'react';
import { BrowserRouter, Route, Routes, useLocation, Navigate } from 'react-router-dom';
import Layout from '../Layout/Layout';
import Dashboard from '../Dashboard/Dashboard';
import Login from '../Login/Login'
import Preferences from '../Preferences/Preferences';
import Accounts from '../Accounts/Accounts.js';
import Friends from '../Friends/Friends.js';
import Transactions from '../Transactions/Transactions';
import Register from '../Register/Register';
import ErrorBoundary from '../ErrorBoundary/ErrorBoundary';
import { useAuth, AuthProvider } from '../Auth/Auth';
import { getAuth, onAuthStateChanged } from "firebase/auth";

function App() {
  const [user, setUser] = useState<string | null>(null);

  const auth = getAuth();
  onAuthStateChanged(auth, (user) => {
    if (user) {
      // User is signed in, see docs for a list of available properties
      // https://firebase.google.com/docs/reference/js/firebase.User
      const email = user.email;
      setUser(email);
      // ...
    } else {
      // User is signed out
      setUser(null);
    }
  });

  // Redirect user to login page if not logged in
  // @ts-ignore
  function RequireAuth({ children }) {
    let auth = useAuth();
    let location = useLocation();

    // @ts-ignore
    if (!user) {
      return <Navigate to="/login" state={{ from: location }} replace />;
    }

    return children
  }

  return (
    
      <AuthProvider>
        <BrowserRouter basename="/">
          <ErrorBoundary>
            <Layout>
              <Routes>
                <Route path="/" element={<h1>Hi, { user }</h1>} />
                <Route path="*" element={<Login />} />
                <Route path="/login" element={<Login />} />
                <Route path="/register" element={<Register />} />
                <Route path="/dashboard" element={<RequireAuth><Dashboard /></RequireAuth>} />
                <Route path="/preferences" element={<RequireAuth><Preferences /></RequireAuth>} />
                <Route path="/accounts" element={<RequireAuth><Accounts /></RequireAuth>} />
                <Route path="/friends" element={<RequireAuth><Friends /></RequireAuth>} />
                <Route path="/transactions" element={<RequireAuth><Transactions /></RequireAuth>} />
              </Routes>
            </Layout>
          </ErrorBoundary>
        </BrowserRouter>
      </AuthProvider>
  );
}

export default App;