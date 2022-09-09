import React, { useState, useEffect, useCallback } from 'react';
import { usePlaidLink } from 'react-plaid-link';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import { getAuth, onAuthStateChanged, User as FirebaseUser } from "firebase/auth";

const PlaidLink = () => {
  const [linkToken, setLinkToken] = useState<string | null>(null);
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

  // Generate link token from server when component mounts
  const generateToken = useCallback(
    async () => {
      await user?.getIdToken()
        .then(async function(jwtToken) {
          const response = await fetch('http://localhost:8080/api/private/linktoken',
        {
            method: 'POST',
            headers: { 
                'Content-Type': 'application/json', 
                'Authorization': `${jwtToken}` 
            },
        }
      )
      if (!response.ok) {
        setLinkToken(null);
        return;
      }
      const data = await response.json();
      if (data) {
        if (data.error != null) {
          setLinkToken(null);
          console.log(data.error);
          return;
        }
        setLinkToken(data.link_token);
      }
      localStorage.setItem("link_token", data.link_token); //to use later for Oauth
        })
    },
    [user]
  );

  useEffect(() => {
    const init = async () => {
      // do not generate a new token for OAuth redirect; instead
      // setLinkToken from localStorage
      if (window.location.href.includes("?oauth_state_id=")) {
        const token = localStorage.getItem("link_token");
        setLinkToken(token);
        return;
      }
      generateToken();
    };
    init();
  }, [generateToken]);

  // Send the generated link public token to the server
  const onSuccess = useCallback(
    (public_token: string) => {
      // send public_token to server
      const setToken = async () => {
        user?.getIdToken()
          .then(async function(jwtToken) {
            const response = await fetch("http://localhost:8080/api/private/accesstoken", {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
                'Authorization': `${jwtToken}`
              },
              body: `public_token=${public_token}`,
            });
            if (!response.ok) {
              return;
            }
            const data = await response.json();
            // dispatch({
            //   type: "SET_STATE",
            //   state: {
            //     itemId: data.item_id,
            //     accessToken: data.access_token,
            //     isItemAccess: true,
            //   },
            // });                
          } 
    )};
      setToken();
      window.history.pushState("", "", "/");
    },
    []
  );

  const config: Parameters<typeof usePlaidLink>[0] = {
    token: linkToken!,
    onSuccess,
  };
  const { open, ready } = usePlaidLink(config);

  return (  
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
  );
};

PlaidLink.displayName = "PlaidLink";

export default PlaidLink;