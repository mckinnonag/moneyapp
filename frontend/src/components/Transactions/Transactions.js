import React, { useState } from 'react';
import { DataGrid, onSelectionModelChange } from '@mui/x-data-grid';
import Button from '@mui/material/Button';
import { useEffect } from 'react';

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
//   {
//     field: 'fullName',
//     headerName: 'Full name',
//     description: 'This column has a value getter and is not sortable.',
//     sortable: false,
//     width: 160,
//     valueGetter: (params) =>
//       `${params.row.firstName || ''} ${params.row.lastName || ''}`,
//   },
];

const Transactions = () => {
  const [rows, setRows] = useState([]);
  const [selectedRows, setSelectedRows] = useState([]);
  
  useEffect(() => {
    const fetchTransactions = async () => {
      const jwtToken = JSON.parse(sessionStorage.getItem('token'))['token'];
      const requestOptions = {
          method: 'GET',
          headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${jwtToken}` },
      };
      const response = await fetch('http://localhost:8080/api/private/gettransactions', requestOptions)
      const transactions = await response.json();
      console.log(typeof transactions.transactions);
      setRows(transactions.transactions);
    };
    fetchTransactions();
  }, []);

  function printRows() {
    console.log(selectedRows);
  }

  let txList = rows.map(function(tx){
    return {id: tx.ID, 
            merchant: tx.MerchantName, 
            amount: tx.Amount,
            date: tx.Date,
          };
  });

  return(
    <div style={{ height: 600, width: '100%' }}>
      <DataGrid
        rows={ txList }
        columns={columns}
        pageSize={15}
        rowsPerPageOptions={[15]}
        checkboxSelection
        onSelectionModelChange={(selectedFile) => {
          setSelectedRows(selectedFile);
        }} 
      />
      <Button variant="outlined" onClick={printRows}>Split selected</Button>
    </div>
  )
};

export default Transactions;