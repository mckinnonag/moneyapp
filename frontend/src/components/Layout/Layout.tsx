import * as React from 'react';
import { useState } from 'react';
import { ThemeProvider } from '@mui/material/styles';
import { CssBaseline, GlobalStyles  } from '@mui/material/';
import Box from '@mui/material/Box';
import Nav from '../Nav/Nav';
import Theme from '../Theme/Theme';
import Typography from '@mui/material/Typography';
import Link from '@mui/material/Link';
import { getAuth, onAuthStateChanged } from "firebase/auth";

const theme = Theme;

// Dynamically polls the current year for copyright data
function Copyright() {
    return (
      <Typography variant="body2" color="text.secondary" align="center">
        {'Copyright © '}
        <Link color="inherit" href="https://mui.com/">
          Your Website
        </Link>{' '}
        {new Date().getFullYear()}
        {'.'}
      </Typography>
    );
  }

function Layout (props: any) {
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

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline/>
      <GlobalStyles styles={{ }}/>
      <main>
        {/* { user && <Nav /> } */}
        <Nav>
          {props.children}
        </Nav>
      </main>
      {/* Footer */}
      <Box sx={{ bgcolor: 'background.paper', p: 6 }} component="footer">
          <Copyright />
      </Box>
      {/* End footer */}
    </ThemeProvider>
  );
}

export default Layout;