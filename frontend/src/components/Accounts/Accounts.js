import React, { useState } from "react";

import PlaidLink from '../PlaidLink/PlaidLink.js'
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';

const bankAccounts = [{
  "type": "Checking",
  "name": "Bank of America",
  "balance": 450,
},
{
  "type": "Savings",
  "name": "Chase",
  "balance": 275,
},]

// let cardsList = bankAccounts.map(function(acct, index){
//   return  <Card key={ index } sx={{ 
//                                     minWidth: 275,
//                                     width: 500,
//                                     margin: "auto",
//                                     "margin-top": 20,
//                                   }}>
//             <CardContent>
//               <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
//                   placeholder
//               </Typography>
//               <Typography variant="h5" component="div">
//                   {acct.name}
//               </Typography>
//               <Typography sx={{ mb: 1.5 }} color="text.secondary">
//                   ${acct.balance}
//               </Typography>
//               <Typography variant="body2">
//                   {acct.type}
//               </Typography>
//             </CardContent>
//             <CardActions>
//               <Button size="small">Remove</Button>
//             </CardActions>
//           </Card>
// })

const Accounts = () => {
  const [accounts, setAccounts] = useState([]);

  React.useEffect(() => {
    fetchAccounts();
  }, []);

  const fetchAccounts = async () => {
    const jwtToken = JSON.parse(sessionStorage.getItem('token'))['token'];
    const requestOptions = {
        method: 'GET',
        headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${jwtToken}` },
    }
    const response = await fetch('http://localhost:8080/api/private/accounts', requestOptions).then(response => 
    response.json().then(data => ({
      data: data
    })).then(res => {
      console.log(res.data)
      setAccounts(res.data.accounts);
      })
    );
  };

  let cardsList = accounts.map(function(acct, index){
    return  <Card key={ index } sx={{ 
                                      minWidth: 275,
                                      width: 500,
                                      margin: "auto",
                                      "margin-top": 20,
                                    }}>
              <CardContent>
                <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
                    {acct.OfficialName}
                </Typography>
                <Typography variant="h5" component="div">
                    {acct.Name}
                </Typography>
                <Typography sx={{ mb: 1.5 }} color="text.secondary">
                    Balance: ${acct.BalanceAvailable}
                </Typography>
                <Typography variant="body2">
                    {acct.Type}
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