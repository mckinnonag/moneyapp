import * as React from 'react';
import { DataGrid } from '@mui/x-data-grid';

const columns = [
  { field: 'id', headerName: 'ID', width: 70 },
  { field: 'vendor', headerName: 'Vendor', width: 130 },
  {
    field: 'amount',
    headerName: 'Amount',
    type: 'number',
    width: 90,
  },
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

const rows = [
    {id: 1, vendor: 'Test1', amount: 41.0, date: "1/21/22"},
    {id: 2, vendor: 'Test2', amount: 14.50, date: "1/22/22"}
];

export default function Transactions() {
  return (
    <div style={{ height: 400, width: '100%' }}>
      <DataGrid
        rows={rows}
        columns={columns}
        pageSize={5}
        rowsPerPageOptions={[5]}
        checkboxSelection
      />
    </div>
  );
}


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