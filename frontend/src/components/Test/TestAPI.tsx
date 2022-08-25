import React, { useState } from 'react';
import Button from '@mui/material/Button';
import { AuthProvider } from '../Auth/Auth';
import { getAuth, onAuthStateChanged, User as FirebaseUser } from "firebase/auth";

export default function Test() {
    const [user, setUser] = useState<FirebaseUser | null>(null);

    const auth = getAuth();
    onAuthStateChanged(auth, (user) => {
        if (user) {
            // User is signed in
            const u = user;
        setUser(u);
        } else {
            // User is signed out
            setUser(null);
        }
    });

    async function test() {
        user?.getIdToken().then((token) => {
            fetch('http://localhost:8080/api/public/test', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ 
                token: token,
                })
            })
        })
        
    }
    
    return (
        <Button 
            variant="outlined"
            onClick={
                test
            }
        >
            Test!</Button>
    )
}
