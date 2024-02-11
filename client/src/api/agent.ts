import type { Transaction } from "@/models/transaction.ts";
import axios, { type AxiosResponse } from "axios";
axios.defaults.baseURL = "http://localhost:8080/api/";
axios.defaults.withCredentials = true;
const responseBody = (response: AxiosResponse) => response.data;
const requests = {
   get: (url: string) => axios.get(url).then(responseBody),
   post: (url: string, body: {}) => axios.post(url, body).then(responseBody),
   put: (url: string, body: {}) => axios.put(url, body).then(responseBody),
   delete: (url: string) => axios.delete(url).then(responseBody),
};
const transaction = {
   get: () => requests.get("transactions"),
   addTransaction: (transaction: Transaction) => {
      return requests.post("addTransaction", transaction);
   },
   deleteTransaction: (id: number) => requests.delete(`deleteTransaction/${id}`),
};

const agent = {
   transaction,
};
export default agent;
