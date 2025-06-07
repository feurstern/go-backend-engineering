import { ref, Ref } from 'vue';

export default () => {
  const leftDrawerOpen: Ref<boolean> = ref(false);
  const rightDrawerOpen: Ref<boolean> = ref(false);

  const toggleLeftDrawer = (): void => {
    leftDrawerOpen.value = !leftDrawerOpen.value;
  };

  const toggleRightDrawer = (): void => {
    rightDrawerOpen.value = !rightDrawerOpen.value;
  };

  return {
    leftDrawerOpen,
    rightDrawerOpen,
    toggleLeftDrawer,
    toggleRightDrawer,
  };
};
