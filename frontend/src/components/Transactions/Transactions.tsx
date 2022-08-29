import React, { useState } from 'react';
import { DataGrid } from '@mui/x-data-grid';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import { useEffect } from 'react';
import { getAuth, onAuthStateChanged, User as FirebaseUser } from "firebase/auth";
import { getTransactions } from './getTransactions';

const columns = [
  {
    field: 'amount',
    headerName: 'Amount',
    type: 'number',
    currencySymbol: 'USD',
    width: 70,
  },
  { field: 'merchant', headerName: 'Merchant', width: 300 },
  { field: 'date', headerName: 'Date', width: 120},
];

const Transactions = () => {
  const [rows, setRows] = useState([]);
  const [selectedRows, setSelectedRows] = useState([]);

  // JWT Token
  const [user, setUser] = useState<FirebaseUser | null>(null);

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
  
  // useEffect(() => {
  //   const fetchTransactions = async () => {
  //     const jwtToken = JSON.parse(sessionStorage.getItem('token'))['token'];
  //     const requestOptions = {
  //         method: 'GET',
  //         headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${jwtToken}` },
  //     };
  //     const response = await fetch('http://localhost:8080/api/private/gettransactions', requestOptions)
  //     const transactions = await response.json();
  //     console.log(typeof transactions.transactions);
  //     setRows(transactions.transactions);
  //   };
  //   fetchTransactions();
  // }, []);

  // Load transaction data from the server
//   useEffect(() => {
//     const fetchTransactions = async () => {
//       user?.getIdToken().then((token) => {
//         fetch('http://localhost:8080/api/private/gettransactions', 
//           { 
//             method: 'GET',
//             headers: {
//               'Content-Type': 'application/json'
//             },
//             // body: JSON.stringify({ 
//             //     token: token,
//             //     })
//           }).then((response) => {
//             // @ts-ignore
//             setRows(response.json().transactions);
//           })
//         })
      
//     //   const transactions = await response.json();
//     //   setRows(transactions.transactions);
//     };
//     fetchTransactions();
//   }, []);

//   useEffect(() => {
//     getTransactions(user)
//       .then(transactions => {
//         setRows(transactions)
//       })
//   }, [])

  function printRows() {
    // To be replaced - this function is being used to test interactions with rows in the datagrid.
    let sharedTransactions = [];
    for (let i = 0; i < selectedRows.length; i++) {
      // @ts-ignore
      const tx = rows.filter(row => row.ID === selectedRows[i]);
      const friend = 'test_friend';
      sharedTransactions.push({
        transaction: tx,
        friend: friend,
      })
    }
  }

  let txList = rows.map(function(tx){
    // @ts-ignore
    return {id: tx.ID, 
        // @ts-ignore
            merchant: tx.MerchantName, 
            // @ts-ignore
            amount: tx.Amount,
            // @ts-ignore
            date: tx.Date,
          };
  });

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
      <DataGrid
        rows={ txList }
        columns={columns}
        pageSize={15}
        rowsPerPageOptions={[15]}
        checkboxSelection
        onSelectionModelChange={(selectedFile) => {
            // @ts-ignore
          setSelectedRows(selectedFile);
        }} 
      />
      <Button variant="outlined" onClick={printRows}>Split selected</Button>
    </Box>
  )
};

export default Transactions;