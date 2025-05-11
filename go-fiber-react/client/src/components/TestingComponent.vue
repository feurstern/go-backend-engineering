<script lang="ts">
import { title } from 'process';
import { Skills } from './models';

interface TestingProps {
  title: string;
  skill: Skills[];
}

const props = withDefaults(defineProps<TestingProps>(), {
  skill: () => [],
});
</script>

<template>
  <p>title {{ title }}</p>
  <ul>
    <li v-for="x in skill" :key="x.id">Language : {{ x }}</li>
  </ul>
</template>
