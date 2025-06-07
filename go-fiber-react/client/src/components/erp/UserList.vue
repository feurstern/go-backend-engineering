<script setup lang="ts">
import useUserList from './UserList';

const { rows, columns } = useUserList();
</script>

<template>
  <div class="q-pa-md">
    <q-table flat bordered :rows="rows" :columns="columns" />
  </div>
</template>
