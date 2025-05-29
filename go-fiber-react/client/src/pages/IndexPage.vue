<script setup lang="ts">
import { Component, defineAsyncComponent, ref } from 'vue';
import type { Todo, Meta, Skills } from 'components/models';
import ExampleComponent from 'components/ExampleComponent.vue';

const todos = ref<Todo[]>([
  {
    id: 1,
    content: 'ct1',
  },
  {
    id: 2,
    content: 'ct2',
  },
  {
    id: 3,
    content: 'ct3',
  },
  {
    id: 4,
    content: 'ct4',
  },
  {
    id: 5,
    content: 'ct5',
  },
]);

const TestingComponent: Component = defineAsyncComponent(
  () => import('components/TestingComponent.vue'),
);

const meta = ref<Meta>({
  totalCount: 1200,
});

const skills = ref<Skills[]>([
  {
    id: 1,
    language: 'Go',
    level: 'Advance',
  },
  {
    id: 2,
    language: 'Rust',
    level: 'Advance',
  },
]);
</script>

<template>
  <q-page class="row items-center justify-evenly">
    <example-component
      title="Example component"
      active
      :todos="todos"
      :meta="meta"
    ></example-component>
    <testing-component title="test" :skill="skills" />
  </q-page>
</template>
