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
          const response = await fetch('http://localhost:8080/v1/api/private/plaid/create_link_token',
        {
            method: 'POST',
            headers: { 
                'Content-Type': 'application/json', 
                'Authorization': `Bearer ${jwtToken}` 
            },
        }
      )
      if (!response.ok) {
        console.log("create link token response not OK")
        setLinkToken(null);
        return;
      }
      const data = await response.json();
      if (data) {
        if (data.error != null) {
          setLinkToken(null);
          console.log("error in response data")
          console.log(data.error);
          return;
        }
        if (data.link_token == null) {
          console.log("empty link token");
          return;
        } 
        setLinkToken(data.link_token);
        localStorage.setItem("link_token", data.link_token); //to use later for Oauth
        
      } else {
        console.log("no data in response")
      }
        })
    },
    [user]
  );

  useEffect(() => {
    const init = async () => {
      // do not generate a new token for OAuth redirect; instead
      // setLinkToken from localStorage
      // if (window.location.href.includes("?oauth_state_id=")) {
      //   const token = localStorage.getItem("link_token");
      //   setLinkToken(token);
      //   return;
      // }
      generateToken();
    };
    init();
  }, [generateToken]);

  // Send the generated link public token to the server
  const onSuccess = useCallback(
    (linkToken: string) => {
      // send public_token to server
      const setToken = async () => {
        user?.getIdToken()
          .then(async function(jwtToken) {
            const data = {public_token: linkToken}
            const response = await fetch("http://localhost:8080/v1/api/private/plaid/set_access_token", {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
                'Authorization': `Bearer ${jwtToken}`
              },
              body: JSON.stringify(data),
            });
            console.log(response);
            if (!response.ok) {
              return;
            }
            // const data = await response.json();
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
    [user, linkToken]
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