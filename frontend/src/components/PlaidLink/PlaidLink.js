import React, { useCallback, useState } from 'react';
import { usePlaidLink, PlaidLinkOnSuccess } from 'react-plaid-link';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';

const PlaidLink = () => {
  const [token, setToken] = useState(null);

  // get link_token from your server when component mounts
  React.useEffect(() => {
    const createLinkToken = async () => {
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${token}` },
            body: JSON.stringify({ title: 'React POST Request Example' })
        }
        const response = await fetch('http://localhost:8080/api/private/linktoken', requestOptions);
        const { link_token } = await response.json();
      setToken(link_token);
    };
    createLinkToken();
  }, [token]);

  const onSuccess = useCallback((publicToken, metadata) => {
    // send public_token to your server
    // https://plaid.com/docs/api/tokens/#token-exchange-flow
    console.log(publicToken, metadata);
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${token}` },
        body: JSON.stringify({ token: publicToken })
    }
    const response = fetch('http://localhost:8080/api/private/accesstoken', requestOptions);
    console.log(response);
  }, [token]);

//   const { open, ready } = usePlaidLink({
//     token,
//     onSuccess,
//     // onEvent
//     // onExit
//   });

    let isOauth = false;
    const config = {
        token,
        onSuccess,
    };

    if (window.location.href.includes("?oauth_state_id=")) {
        // TODO: figure out how to delete this ts-ignore
        // @ts-ignore
        config.receivedRedirectUri = window.location.href;
        isOauth = true;
    }

    const { open, ready } = usePlaidLink(config);

    // useEffect(() => {
    //     if (isOauth && ready) {
    //       open();
    //     }
    //   }, [ready, open, isOauth]);

    return (
      <div align="center">
        <Box 
          sx={{
            mx: "auto",
            m: 1,
          }}
          >
          <Button 
            variant="outlined"
            onClick={() => open()} disabled={!ready}>
            Connect a bank account
          </Button>
        </Box>
      </div>
    );
};

export default PlaidLink;