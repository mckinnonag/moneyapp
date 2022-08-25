import React, { useState } from 'react';
import { BrowserRouter, Route, Routes, useLocation, Navigate } from 'react-router-dom';
import Layout from '../Layout/Layout';
import Dashboard from '../Dashboard/Dashboard';
import Login from '../Auth/Login';
import Logout from '../Auth/Logout';
import Preferences from '../Preferences/Preferences';
import Accounts from '../Accounts/Accounts';
import Friends from '../Friends/Friends';
import Transactions from '../Transactions/Transactions';
import Register from '../Auth/Register';
import ErrorBoundary from '../ErrorBoundary/ErrorBoundary';
import { AuthProvider } from '../Auth/Auth';
import { getAuth, onAuthStateChanged } from "firebase/auth";

function App() {
  const [user, setUser] = useState<string | null>(null);

  const auth = getAuth();
  onAuthStateChanged(auth, (user) => {
    if (user) {
      // User is signed in
      const uid = user.uid;
      setUser(uid);
    } else {
      // User is signed out
      setUser(null);
    }
  });

  // Redirect user to login page if not logged in
  // @ts-ignore
  function RequireAuth({ children }) {
    let location = useLocation();

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
                <Route path="/" element={<RequireAuth><Dashboard /></RequireAuth>} />
                <Route path="*" element={<Login />} />
                <Route path="/login" element={<Login />} />
                <Route path="/logout" element={<Logout />} />
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