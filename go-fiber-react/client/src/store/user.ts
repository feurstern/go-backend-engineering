import { defineStore } from 'pinia';
import { ref } from 'vue';

export const useUserStore = defineStore('user', () => {
  const user = ref([]);

  const SET_USER_LIST = (data: []) => {
    user.value = data;
  };
});
