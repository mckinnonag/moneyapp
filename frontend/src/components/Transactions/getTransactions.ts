import { User as FirebaseUser } from "firebase/auth";

export async function getTransactions(user: FirebaseUser | null) {
    // return fetch('http://localhost:8080/api/private/gettransactions')
    //   .then(data => data.json())
  //   return await user?.getIdToken()
  //       .then(async function(jwtToken) {
  //         const response = await fetch('http://localhost:8080/api/private/gettransactions',
  //       {
  //           method: 'GET',
  //           headers: { 
  //               'Content-Type': 'application/json', 
  //               'Authorization': `${jwtToken}` 
  //           },
  //       }
  //     )
  //   }).then(data => data.json())
  // }
  const response = await user?.getIdToken().then((token) => {
    fetch('http://localhost:8080/api/private/gettransactions',
        {
            method: 'GET',
            headers: { 
                'Content-Type': 'application/json', 
                'Authorization': `${token}` 
            },
        }
      ).then(data => data.json())
  });
  return response;
}