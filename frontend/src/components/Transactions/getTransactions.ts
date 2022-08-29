import { User as FirebaseUser } from "firebase/auth";

export function getTransactions(user: FirebaseUser | null) {
    return fetch('http://localhost:8080/api/private/gettransactions')
      .then(data => data.json())
  }