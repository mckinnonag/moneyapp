import React, { useState } from "react";
import PlaidLink from '../PlaidLink/PlaidLink'
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import { getAuth, onAuthStateChanged, User as FirebaseUser } from "firebase/auth";

const Accounts = () => {
  const [accounts, setAccounts] = useState([]);
  const [user, setUser] = useState<FirebaseUser | null>(null);  // JWT Token

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

  React.useEffect(() => {
    fetchAccounts();
  }, []);

  const fetchAccounts = async () => {
    user?.getIdToken().then((token) => {
      fetch('http://localhost:8080/api/private/accounts', 
      {
        method: 'GET',
        headers: { 
          'Content-Type': 'application/json', 
          'Authorization': `Bearer ${token}` 
        },
      }
      ).then((accts) => {
        // @ts-ignore
        setAccounts(accts.accounts);
      })
  })};

  let cardsList = accounts.map(function(acct, index){
    return  <Card 
                key={ index } 
                sx={{ 
                    minWidth: 275,
                    width: 500,
                    margin: "auto",
                    "margin-top": 20,
                }}
            >
              <CardContent>
                <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
                    {/* {acct.OfficialName} */}
                </Typography>
                <Typography variant="h5" component="div">
                    {/* {acct.Name} */}
                </Typography>
                <Typography sx={{ mb: 1.5 }} color="text.secondary">
                    {/* Balance: ${acct.BalanceAvailable} */}
                </Typography>
                <Typography variant="body2">
                    {/* {acct.Type} */}
                </Typography>
              </CardContent>
              <CardActions>
                <Button size="small">Remove</Button>
              </CardActions>
            </Card>
  })

  return(
    <div>
        { cardsList }
        <PlaidLink /> 
    </div>
  );
}

export default Accounts;