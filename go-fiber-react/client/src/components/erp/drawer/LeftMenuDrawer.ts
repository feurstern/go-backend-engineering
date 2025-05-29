import { Ref, ref } from 'vue';

export default () => {
  const menu: Ref<string> = ref('inbox');

  return {
    menu,
  };
};
