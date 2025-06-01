import { useQuasar } from 'quasar';
import { userRegistrationPayload } from 'src/components/models';
import { userRegistration } from 'src/service/user';
import { ref, Ref } from 'vue';

export default () => {
  const username: Ref<string> = ref('');
  const email: Ref<string> = ref('');
  const password: Ref<string> = ref('');
  const isPwd: Ref<boolean> = ref(false);

  const handleSubmit = async () => {
    const $q = useQuasar();
    try {
      const payload: userRegistrationPayload = {
        username: username.value,
        email: email.value,
        password: password.value,
      };

      const res = await userRegistration(payload);

      if (res.success) {
        $q.notify('succes');
      }
    } catch (error) {
      console.log('error', error);
      $q.notify('errrr');
    }
  };

  return {
    username,
    handleSubmit,
    email,
    password,
    isPwd,
  };
};
