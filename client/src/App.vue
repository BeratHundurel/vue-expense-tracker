<script setup lang="ts">
import Header from "./components/MainHeader.vue";
import IncomeExpenses from "./components/IncomeExpenses.vue";
import TransactionList from "./components/TransactionList.vue";
import AddTransaction from "./components/AddTransaction.vue";
import Balance from "./components/BalanceList.vue";
import { computed, ref, onMounted } from "vue";
import { type Transaction } from "./models/transaction.ts";
import { useToast } from "vue-toastification";
import agent from "./api/agent";
const toast = useToast();
const transactions = ref<Transaction[]>([]);

onMounted(() => {
   const localTransactions = localStorage.getItem("transactions");
   console.log("localTransactions", localTransactions);
   if (localTransactions === null) {
      console.log("fetching transactions");
      agent.transaction.get().then((result) => {
         transactions.value = result;
         saveTransactionsToLocalStorage();
      });
   } else {
      console.log("loading transactions from local storage");
      transactions.value = JSON.parse(localTransactions);
   }
});
const total = computed(() => {
   return transactions.value.length > 0 ? transactions.value.reduce((acc, item) => (acc += item.amount), 0) : 0;
});

const income = computed(() => {
   return transactions.value
      .filter((transaction) => transaction.amount > 0)
      .reduce((acc, transaction) => acc + transaction.amount, 0)
      .toFixed(2);
});

const expense = computed(() => {
   return transactions.value
      .filter((transaction) => transaction.amount < 0)
      .reduce((acc, transaction) => acc + transaction.amount, 0)
      .toFixed(2);
});

const handleTransactionSubmitted = (transaction: Transaction) => {
   agent.transaction.addTransaction(transaction);
   transactions.value.push(transaction);
   saveTransactionsToLocalStorage();
};

const deleteTransaction = (id: number) => {
   agent.transaction.deleteTransaction(id);
   transactions.value = transactions.value.filter((transaction) => transaction.id !== id);
   saveTransactionsToLocalStorage();
};

const saveTransactionsToLocalStorage = () => {
   localStorage.setItem("transactions", JSON.stringify(transactions.value));
};
</script>

<template>
   <Header />
   <div class="container">
      <Balance :total="total" />
      <IncomeExpenses :income="parseFloat(income)" :expense="parseFloat(expense)" />
      <TransactionList :transactions="transactions" @deleteTransaction="deleteTransaction" />
      <AddTransaction @addTransaction="handleTransactionSubmitted" />
   </div>
</template>
