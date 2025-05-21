<script setup lang="ts">
import { Skills } from './models';
import { ref, Ref, defineProps } from 'vue';

// const t1: Ref<string> = ref('');
interface TestingProps {
  title: string;
  skill: Skills[];
}

const props = defineProps<TestingProps>();
</script>
<template>
  <p>title: {{ props.title }}</p>
  <ul>
    <li class="bg-red text-white" v-for="x in props.skill" :key="x.id">
      Language: {{ x.language }}
    </li>
  </ul>
</template>
