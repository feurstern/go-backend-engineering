import { useQuasar } from 'quasar';
import { login } from 'src/service/user';
import { ref, Ref } from 'vue';

export default () => {
  const email: Ref<string> = ref('');
  const password: Ref<string> = ref('');
  const $q = useQuasar();

  const handleLogin = async () => {
    const paylaod = {
      email: email.value,
      password: password.value,
    };
    try {
      const res = await login(paylaod);

      if (res.success) {
        $q.notify({
          message: res.message,
          icon: 'check_circle',
          color: 'positive',
        });
      } else {
        $q.notify({
          message: res.message,
          icon: 'error',
          color: 'negative',
        });
      }
    } catch (error) {
      $q.notify({
        message: 'error',
        icon: 'error',
        color: 'negative',
      });
    }
  };

  const onReset = (): void => {
    email.value = '';
    password.value = '';
  };
  return {
    handleLogin,
    onReset,
    email,
    password,
  };
};
