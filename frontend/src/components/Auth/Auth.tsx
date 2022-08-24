import React, {useContext, useState, useEffect } from 'react';
// import firebase from "firebase/app";
import { User as FirebaseUser } from "firebase/auth";
import { Auth } from '../../firebase';

const AuthContext = React.createContext<FirebaseUser | null>(null);

export const AuthProvider: React.FC = ({ children }) => {
  const [user, setUser] = useState<FirebaseUser | null>(null);

  // 
  useEffect(() => {
    const unsubscribe = Auth.onAuthStateChanged((firebaseUser) => {
      setUser(firebaseUser);
    });

    return unsubscribe;
  }, []);

  return <AuthContext.Provider value={user}>{children}</AuthContext.Provider>;
};

export function useAuth(){
  return useContext(AuthContext)
}