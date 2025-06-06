import { defineStore } from 'pinia';
import { Users } from 'src/components/models';
import { Ref, ref } from 'vue';

export const useUserStore = defineStore('user', () => {
  const user = ref<Users>();

  const SET_USER_LIST = (data: Users) => {
    user.value = data;
  };
});
