import * as React from 'react';
import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { getAuth, signOut } from "firebase/auth";
import CircularProgress from '@mui/material/CircularProgress';

export default function Logout() {
    const auth = getAuth();
    const navigate = useNavigate();

    useEffect(() => {
        signOut(auth).then(() => {
            // Sign-out successful.
            navigate("/", { replace: true });
        }).catch((error) => {
            // An error happened.
            console.log(error);
        });
    
    }, []);

    return (
        <CircularProgress />
    );
}