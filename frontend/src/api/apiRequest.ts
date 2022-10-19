import { User as FirebaseUser } from "firebase/auth";
import { base_url } from "./appURLs";

export function Client(user: FirebaseUser | null, endpoint: string, method: string) {
    return user?.getIdToken().then(async (token) => {
        const headers = {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
        }
        const config = {
            method: method,
            headers: {
            ...headers,
            },
            // body: body
        }

      const response = await fetch(`${base_url}${endpoint}`, config)
      if (response.ok) {
        console.log(response.ok);
        return response.json();
      } else {
        const errorMessage = await response.text();
        return Promise.reject(new Error(errorMessage));
      }
      })

      // fetch(`${base_url}${endpoint}`, config)
      // .then(async response => {
      //   if (response.ok) {
      //     return await response.json();
      //   } else {
      //     const errorMessage = await response.text();
      //     return Promise.reject(new Error(errorMessage));
      //   }
      // })
  // })
}
