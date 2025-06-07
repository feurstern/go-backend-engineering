import { computed, WritableComputedRef } from 'vue';

export interface ErpPageProps {
  menu: string;
  title: string;
}

export interface ErpPageEmits {
  (e: 'get:menu', value: string): void;
  (e: 'get:title', value: string): void;
}

export default (props: ErpPageProps, emits: ErpPageEmits) => {
  const title: WritableComputedRef<string> = computed({
    get: () => {
      return props.title;
    },
    set: (value: string) => {
      emits('get:title', value);
    },
  });

  const menu: WritableComputedRef<string> = computed({
    get: () => {
      return props.menu;
    },
    set: (value: string) => {
      emits('get:menu', value);
    },
  });
  return {
    title,
    menu,
  };
};
