<script setup lang="ts">
import { Component, defineAsyncComponent } from 'vue';

const userTableList: Component = defineAsyncComponent(() => import('components/erp/UserList.vue'));
</script>
<template>
  <user-table-list />
</template>
