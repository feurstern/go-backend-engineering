<script setup lang="ts">
import { handleError, onMounted } from 'vue';
import useUserList from './UserList';

const { rows2, columns, fetchUserList } = useUserList();

console.log('rows2', rows2);
onMounted(() => {
  rows2;
  fetchUserList()
    .catch((err) => console.log(err))
    .then();
});
</script>

<template>
  <div class="q-pa-md">
    <q-table flat bordered :rows="rows2" :columns="columns"> </q-table>
  </div>
</template>
