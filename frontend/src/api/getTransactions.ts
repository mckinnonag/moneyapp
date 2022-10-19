import { User as FirebaseUser } from "firebase/auth";
import { Client } from "../api/apiRequest";

export async function getTransactions(user: FirebaseUser | null) {
  const endpoint = '/v1/api/private/plaid/transactions';
  const method = 'GET';

  const today = Date.now();
  const dateOffset = (24*60*60*1000);
  const start = (today - dateOffset).toLocaleString();
  const end = today.toLocaleString();
  const count = 100;
  const offset = 0;

  return Client(user, endpoint, method);
}