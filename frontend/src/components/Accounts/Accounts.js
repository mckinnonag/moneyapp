import React from "react";

import PlaidLink from '../PlaidLink/PlaidLink.js'
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Stack from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import { Container } from "@mui/material";
import { styled } from '@mui/material/styles';

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
  ...theme.typography.body2,
  padding: theme.spacing(1),
  textAlign: 'center',
  color: theme.palette.text.secondary,
}));

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

const cardsList = bankAccounts.map(function(acct, index){
  return  <Card key={ index } sx={{ 
                                    minWidth: 275,
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
                  ${acct.balance}
              </Typography>
              <Typography variant="body2">
                  {acct.type}
              </Typography>
            </CardContent>
            <CardActions>
              <Button size="small">Remove</Button>
            </CardActions>
          </Card>
})

export default function Accounts() {
    return(
      <div>
          { cardsList }
          <PlaidLink /> 
      </div>
    );
  }