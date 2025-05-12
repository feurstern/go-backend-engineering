<script setup lang="ts">
import { Component, defineAsyncComponent } from 'vue';

const CarouselComponent: Component = defineAsyncComponent(() => import('components/Carousel.vue'));
</script>

<template>
  <q-section>
    <q-card> this where the testing page will be rendered </q-card>
    <q-card-section>
      <carousel-component />
    </q-card-section>
  </q-section>
</template>
