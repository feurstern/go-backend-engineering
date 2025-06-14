<script lang="ts" setup>
import { Component, defineAsyncComponent } from 'vue';

const loginComponent: Component = defineAsyncComponent(
  () => import('components/LoginComponent.vue'),
);
</script>

<template>
  <login-component />
</template>
