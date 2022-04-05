import React, { useState } from 'react';
import { DataGrid } from '@mui/x-data-grid';
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

  let txList = rows.map(function(tx, index){
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
      />
    </div>
  )
};

export default Transactions;

// export default function Transactions() {
//   return (
    // <div style={{ height: 400, width: '100%' }}>
    //   <DataGrid
    //     rows={rows}
    //     columns={columns}
    //     pageSize={5}
    //     rowsPerPageOptions={[5]}
    //     checkboxSelection
    //   />
    // </div>
//   );
// }


// class Transactions extends Component {
//     constructor(props) {
//         super(props)
//         this.state = {
//             transactions: [
//                 {id: 1, vendor: 'Test1', amount: 41.0, date: "1/21/22"},
//                 {id: 2, vendor: 'Test2', amount: 14.50, date: "1/22/22"}
//             ]
//         }
//     }

//     renderTableData() {
//         return this.state.transactions.map((transaction, index) => {
//             const {id, vendor, amount, date} = transaction
//             return (
//                 <tr key={id}>
//                     <td>{id}</td>
//                     <td>{vendor}</td>
//                     <td>${amount}</td>
//                     <td>{date}</td>
//                 </tr>
//             )
//         })
//     }

//     renderTableHeader() {
//         let header = Object.keys(this.state.transactions[0])
//         return header.map((key, index) => {
//             return <th key={index}>{key.toUpperCase()}  </th>
//         })
//     }

//     render() {
//         return (
//             <>
//                 <table id='transactions'>
//                     <tbody>
//                         <th>{this.renderTableHeader()}</th>
//                         {this.renderTableData()}
//                     </tbody>
//                 </table>
//             </>
//         )
//     }
// }

// export default Transactions