import { defineStore } from 'pinia';
import { ref, Ref } from 'vue';

export const useERPStore = defineStore('erp', () => {
  const selectedMenu: Ref<string> = ref('finance');
  const index: Ref<number> = ref(0);
  const SET_SELECTED_MENU = (s: string) => {
    selectedMenu.value = s;
  };
  const SET_SELECTED_INDEX = (v: number) => {
    index.value = v;
  };

  return { selectedMenu, SET_SELECTED_INDEX, SET_SELECTED_MENU };
});
