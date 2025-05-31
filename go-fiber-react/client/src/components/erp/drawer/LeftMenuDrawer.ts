import { Ref, ref } from 'vue';
import { useERPStore } from 'src/store/erp';
import { storeToRefs } from 'pinia';
export default () => {
  const menu: Ref<string> = ref('inbox');
  
  const erp = useERPStore();
  const { selectedMenu } = storeToRefs(erp);
  const { SET_SELECTED_MENU } = erp;

  const handleMenuChange = (m: string): void => {
    menu.value = m;
    SET_SELECTED_MENU(m);
    console.log('seelcted menu', selectedMenu.value);
  };

  return {
    menu,
    erp,
    handleMenuChange,
    selectedMenu,
    SET_SELECTED_MENU,
  };
};
