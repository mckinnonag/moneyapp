import React, { useState } from 'react';
import { DataGrid } from '@mui/x-data-grid';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import { useEffect } from 'react';
import { getAuth, onAuthStateChanged, User as FirebaseUser } from "firebase/auth";
import { getTransactions } from '../../api/getTransactions';
import { Transaction, ConvertTransaction } from '../../types/transaction';

const columns = [
  {
    field: 'amount',
    headerName: 'Amount',
    type: 'number',
    currencySymbol: 'USD',
    width: 70,
  },
  { field: 'merchant_name', headerName: 'Merchant', width: 300 },
  { field: 'category', headerName: 'Category', width: 300},
  { field: 'date', headerName: 'Date', width: 120},
  { field: 'shared with', headerName: 'Shared with', width: 200},
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

  useEffect(() => {
    getTransactions(user)
      .then((data) => {
        setRows(data.transactions);
        console.log(data.transactions[0])
      })
  }, [user])

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
    console.log(sharedTransactions);
  }

  let txList = rows.map(function(tx){
    return ConvertTransaction(tx);
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