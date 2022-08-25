import * as React from 'react';
import { useState } from 'react';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import Alert from '@mui/material/Alert';
import { createUserWithEmailAndPassword } from "firebase/auth";
import { Auth } from '../../firebase';

export default function Register() {
  // Firebase logic
  function register(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const data = new FormData(e.currentTarget);
    createUserWithEmailAndPassword(
        Auth,
        data.get('email') as string, 
        data.get('password') as string,
    ).then((userCredential) => {
        // Signed in 
        const user = userCredential.user;
        console.log(user);
        // ...
      })
      .catch((error) => {
        console.log('Error in registration. Try again with a different username or password.')
      });
  }

  // Verify form fields
  const [input, setInput] = useState({
    email: '',
    password: '',
    retypepassword: ''
  });
 
  const [error, setError] = useState('');
 
  const onInputChange = (
    e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>
    ) => {
    const { name, value } = e.currentTarget;
    setInput(prev => ({
      ...prev,
      [name]: value
    }));
    validateInput(e);
  }
 
  const validateInput = (
    e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>
    ) => {
    const { name, value } = e.currentTarget;   
      switch (name) {
        case "email":
          if (!value) {
            setError("Please enter your email.");
          }
          break;
   
        case "password":
          if (!value) {
            setError("Please enter your password.");
          }
          break;

        case "retypepassword":
          if (!value) {
            setError("Passwords don't match.");
          }
          break;
   
        default:
          break;
      }
    };

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <Box
        sx={{
          marginTop: 8,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign up
        </Typography>
        <Box component="form" onSubmit={register} noValidate sx={{ mt: 1 }}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            autoFocus
            onChange={onInputChange}
            onBlur={validateInput}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            onChange={onInputChange}
            onBlur={validateInput}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="retypepassword"
            label="Retype Password"
            type="password"
            id="retypepassword"
            autoComplete="current-password"
            onChange={onInputChange}
            onBlur={validateInput}
          />
          { error && <Alert severity="error">{error}</Alert> }
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            Register
          </Button>
        </Box>
      </Box>
    </Container>
  );
}