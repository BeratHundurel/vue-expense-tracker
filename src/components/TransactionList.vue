<template>
   <h3>History</h3>
   <ul id="list" class="list">
      <li v-for="transaction in transactions" :key="transaction.id" :class="transaction.amount < 0 ? 'minus' : 'plus'">
         {{ transaction.text }}
         <span>${{ transaction.amount }}</span
         ><button @click="deleteTransaction(transaction.id)" class="delete-btn">x</button>
      </li>
   </ul>
</template>
<script setup lang="ts">
import type { Transaction } from "@/models/Transaction.vue";
import { defineProps, defineEmits } from "vue";
const emit = defineEmits(["deleteTransaction"]);
defineProps({
   transactions: {
      type: Array<Transaction>,
      required: true,
   },
});
const deleteTransaction = (id: number) => {
   emit("deleteTransaction", id);
};
</script>
