import React from "react";
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';

const friendAccounts = [{
  "name": "Tiana",
  "owed": 275,
},
{
  "name": "Alex",
  "owed": -450,
},]

const cardsList = friendAccounts.map(function(acct, index){
    let bal = '';
    if (acct.owed >= 0) {
        bal = `${acct.name} owes you $${acct.owed}`
    } else {
        bal = `You owe ${acct.name} $${acct.owed * -1}`;
    }
    return (
    <Grid item xs={12} md={8} lg={9}>
      <Card 
        key={ index } 
        sx={{ minWidth: 275,
              width: 500,
              margin: "auto",
              "margin-top": 20,
              }}
        >
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
      </Grid>
    )
  })

export default function Accounts() {
    return(
      <Box
        component="main"
        sx={{
        backgroundColor: (theme) =>
            theme.palette.mode === 'light'
            ? theme.palette.grey[100]
            : theme.palette.grey[900],
        flexGrow: 1,
        height: '100vh',
        overflow: 'auto',
        }}
    >
      <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
          <Grid container spacing={3}>
            { cardsList }
          </Grid>
      </Container>
    </Box>
    );
  }