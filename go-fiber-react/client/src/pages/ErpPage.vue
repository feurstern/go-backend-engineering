<script setup lang="ts">
import useErpPage from './ErpPage';
import { ErpPageProps, ErpPageEmits } from './ErpPage';
const props = defineProps<ErpPageProps>();
const emits = defineEmits<ErpPageEmits>();

const { title, menu } = useErpPage(props, emits);
</script>

<template>
  <q-page>
    {{ title }} menu {{ menu }}
    <q-card>
      <q-card-section> Finace </q-card-section>
      <q-card-section> Makerting </q-card-section>
    </q-card>
  </q-page>
</template>
