import { defineStore } from 'pinia';
import { ref, Ref } from 'vue';

export const useCountStore = defineStore('counter', () => {
  const count: Ref<number> = ref(0);

  const increment = () => {
    count.value++;
  };

  const decrement = () => {
    count.value--;
  };

  return { count, increment, decrement };
});

export const useERPStore = defineStore('erp', () => {
  const selectedMenu: Ref<string> = ref('');
  const index: Ref<number> = ref(0);
  const SET_SELECTED_MENU = (s: string) => {
    selectedMenu.value = s;
  };
  const SET_SELECTED_INDEX = (v: number) => {
    index.value = v;
  };

  return { selectedMenu, SET_SELECTED_INDEX, SET_SELECTED_MENU };
});
