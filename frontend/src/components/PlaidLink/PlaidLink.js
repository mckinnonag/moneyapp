import React, { useCallback, useState } from 'react';

import { usePlaidLink, PlaidLinkOnSuccess } from 'react-plaid-link';

const PlaidLink = () => {
  const [token, setToken] = useState(null);

  // get link_token from your server when component mounts
  React.useEffect(() => {
    const createLinkToken = async () => {
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ title: 'React POST Request Example' })
        }
        const response = await fetch('http://localhost:8080/plaid/create', requestOptions);
        const { link_token } = await response.json();
      setToken(link_token);
    };
    createLinkToken();
  }, []);

  const onSuccess = useCallback((publicToken, metadata) => {
    // send public_token to your server
    // https://plaid.com/docs/api/tokens/#token-exchange-flow
    console.log(publicToken, metadata);
  }, []);

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
        <button onClick={() => open()} disabled={!ready}>
        Connect a bank account
        </button>
    );
};

export default PlaidLink;