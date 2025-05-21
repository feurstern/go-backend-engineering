<script setup lang="ts">
import { Ref, ref } from 'vue';

const email: Ref<string> = ref('');
const password: Ref<string> = ref('');
</script>

<template>
  <q-page>
    <q-card>
      <q-card-section> Email :<q-input type="text" v-model="email" /> </q-card-section>
      <q-card-section> Password :<q-input type="password" v-model="password" /> </q-card-section>
    </q-card>
  </q-page>
</template>
