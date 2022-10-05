import { User as FirebaseUser } from "firebase/auth";
import { base_url } from "./appURLs";

export function Client(user: FirebaseUser | null, endpoint: string, method: string) {
    return user?.getIdToken().then((token) => {
        const headers = {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
        }
        const config = {
            method: method,
            headers: {
            ...headers,
            },
        }

      fetch(`${base_url}/${endpoint}`, config)
      .then(async response => {
        if (response.ok) {
          return await response.json()
        } else {
          const errorMessage = await response.text()
          return Promise.reject(new Error(errorMessage))
        }
      })
  })
}
