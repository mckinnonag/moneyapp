import * as React from 'react';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom'
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { Auth } from '../../firebase';
import { signInWithEmailAndPassword } from "firebase/auth";
import Alert from '@mui/material/Alert';

export default function Login() {
  // Verify form fields
  const [input, setInput] = useState({
    email: '',
    password: '',
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
   
        default:
          break;
      }
    };
  
  const navigate = useNavigate();

  const signIn = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    
    // Redirect from browser history if the user was redirected here
    const from = '/';

    // Data from login form
    const data = new FormData(e.currentTarget);
    signInWithEmailAndPassword(
      Auth,
      data.get('email') as string, 
      data.get('password') as string,
    )
    .then((userCredential) => {
      // Signed in 
      navigate(from, { replace: true });
    })
    .catch((error) => {
      setError('Invalid username or password');
    });
  };

  return (
    <Container maxWidth="xs">
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
          Sign in
        </Typography>
        <Box component="form" onSubmit={signIn} noValidate sx={{ mt: 1 }}>
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
          { error && <Alert severity="error">{error}</Alert> }
          <FormControlLabel
            control={<Checkbox value="remember" color="primary" />}
            label="Remember me"
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            Sign In
          </Button>
          <Grid container>
            <Grid item xs>
              <Link href="#" variant="body2">
                Forgot password?
              </Link>
            </Grid>
            <Grid item>
              <Link href="/register" variant="body2">
                {"Don't have an account? Sign Up"}
              </Link>
            </Grid>
          </Grid>
        </Box>
      </Box>
    </Container>
  );
}