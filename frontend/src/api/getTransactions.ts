import { User as FirebaseUser } from "firebase/auth";
import { Client } from "../api/apiRequest";

export async function getTransactions(user: FirebaseUser | null) {
  const endpoint = 'api/private/gettransactions';
  const method = 'GET';
  let res = await Client(user, endpoint, method);
  return res;

//   const response = await user?.getIdToken().then((token) => {
//     fetch('http://localhost:8080/api/private/gettransactions',
//         {
//             method: 'GET',
//             headers: { 
//                 'Content-Type': 'application/json', 
//                 'Authorization': `${token}` 
//             },
//         }
//       ).then(data => data.json())
//   });
//   return response;
}