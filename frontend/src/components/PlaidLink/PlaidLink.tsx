import React, { useCallback, useState } from 'react';
import { usePlaidLink, PlaidLinkOnSuccess } from 'react-plaid-link';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import { getAuth, onAuthStateChanged, User as FirebaseUser } from "firebase/auth";

const PlaidLink = () => {
  const [token, setToken] = useState(null);
  const [user, setUser] = useState<FirebaseUser | null>(null); // JWT Token

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

  // get link_token from server when component mounts
  React.useEffect(() => {
    const createLinkToken = async () => {
        user?.getIdToken().then((token) => {
            console.log('Fetching public token from server');
            fetch('http://localhost:8080/api/private/linktoken',
            {
                method: 'POST',
                headers: { 
                    'Content-Type': 'application/json', 
                    'Authorization': `${token}` 
                },
                body: JSON.stringify({ title: 'Placeholder' })
            }
            ).then((link_token) => {
                // @ts-ignore
                setToken(link_token);
                console.log(link_token);    
            })
        });
    }
    createLinkToken();
  }, []);

  const onSuccess = useCallback((publicToken, metadata) => {
    user?.getIdToken().then((token) => {
        console.log('getting access token');
        fetch('http://localhost:8080/api/private/accesstoken',
        {
            method: 'POST',
            headers: { 
                'Content-Type': 'application/json', 
                'Authorization': `${token}` 
            },
            body: JSON.stringify({ token: publicToken })
        }
        ).then((response) => {
            console.log(response);
        })
    });
  }, []);

  const { open, ready } = usePlaidLink({
    token,
    onSuccess,
    // onEvent
    // onExit
  });

    let isOauth = false;

    if (window.location.href.includes("?oauth_state_id=")) {
        // @ts-ignore
        config.receivedRedirectUri = window.location.href;
        isOauth = true;
    }

    // useEffect(() => {
    //     if (isOauth && ready) {
    //       open();
    //     }
    //   }, [ready, open, isOauth]);

    return (
      <div>
        <Box 
          sx={{
            mx: "auto",
            m: 1,
          }}
          >
          <Button 
            variant="outlined"
            onClick={() => open()} 
            disabled={!ready}
          >
            Connect a bank account
          </Button>
        </Box>
      </div>
    );
};

export default PlaidLink;