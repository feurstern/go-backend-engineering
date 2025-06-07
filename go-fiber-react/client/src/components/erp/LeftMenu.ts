import { ref, Ref } from 'vue';
import { useRouter } from 'vue-router';

export default () => {
  const router = useRouter();
  const selectedLeftMenu: Ref<string> = ref('finance');

  
  const handleChange = (menu: string): void => {
    selectedLeftMenu.value = menu;
    router.push(`${menu.toLocaleLowerCase()}`);
  };

  return {
    handleChange,
    selectedLeftMenu,
  };
};
