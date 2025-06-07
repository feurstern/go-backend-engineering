import { useQuasar } from 'quasar';
import { userRegistrationPayload } from 'src/components/models';
import { userRegistration } from 'src/service/user';
import { ref, Ref } from 'vue';

export default () => {
  const $q = useQuasar(); 
  const username: Ref<string> = ref('');
  const email: Ref<string> = ref('');
  const password: Ref<string> = ref('');
  const isPwd: Ref<boolean> = ref(false);

  const handleSubmit = async () => {
    try {
      const payload: userRegistrationPayload = {
        username: username.value,
        email: email.value,
        password: password.value,
      };

      const res = await userRegistration(payload);

      if (res.success) {
        $q.notify({
          message: 'The data has been submitted successfully',
          icon: 'check_circle',
          color: 'positive',
        });
      }
    } catch (error) {
      console.log('error', error);
      $q.notify({
        message: 'Failed to submit data.',
        icon: 'error',
        color: 'negative',
      });
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
