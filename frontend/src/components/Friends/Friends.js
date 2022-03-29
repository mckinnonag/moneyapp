import React from "react";

import PlaidLink from '../PlaidLink/PlaidLink.js'
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';

const friendAccounts = [{
  "name": "Tiana",
  "owed": 275,
},
{
  "name": "Alex",
  "owed": -450,
},]

const cardsList = friendAccounts.map(function(acct, index){
    var bal = '';
    if (acct.owed >= 0) {
        bal = `${acct.name} owes you $${acct.owed}`
    } else {
        bal = `You owe ${acct.name} $${acct.owed * -1}`;
    }
    return <Card key={ index } sx={{ minWidth: 275,
                                    width: 500,
                                    margin: "auto",
                                    "margin-top": 20,
                                    }}>
            <CardContent>
              <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
                  placeholder
              </Typography>
              <Typography variant="h5" component="div">
                  {acct.name}
              </Typography>
              <Typography sx={{ mb: 1.5 }} color="text.secondary">
                  placeholder
              </Typography>
              <Typography variant="body2">
                  { bal }
              </Typography>
            </CardContent>
            <CardActions>
              <Button size="small">Remove Friend  </Button>
            </CardActions>
          </Card>
  })

export default function Accounts() {
    return(
      <div>
        { cardsList }
      </div>
    );
  }