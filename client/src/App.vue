<script setup lang="ts">
import Header from "./components/MainHeader.vue";
import IncomeExpenses from "./components/IncomeExpenses.vue";
import TransactionList from "./components/TransactionList.vue";
import AddTransaction from "./components/AddTransaction.vue";
import Balance from "./components/BalanceList.vue";
import { computed, ref, onMounted } from "vue";
import { type Transaction } from "./models/Transaction.vue";
import { useToast } from "vue-toastification";
import agent from "./app/api/agent.ts";
const toast = useToast();
const transactions = ref<Transaction[]>([]);

onMounted(() => {
   const savedTransactions = JSON.parse(localStorage.getItem("transactions") || "[]");
   console.log(savedTransactions);
   if (savedTransactions) {
      transactions.value = savedTransactions;
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
   agent.transactions.get().then((result) => {
      console.log(result);
      console.log(transaction)
   });
   //    transactions.value.push(transaction);
   //    saveTransactionsToLocalStorage();
   //    toast.success("Transaction added successfully");
};
const deleteTransaction = (id: number) => {
   const index = transactions.value.findIndex((transaction) => transaction.id === id);
   if (index === -1) toast.error("Transaction not found");
   transactions.value.splice(index, 1);
   saveTransactionsToLocalStorage();
   toast.success("Transaction deleted successfully");
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
