import { Ref, ref } from 'vue';
import { getUserList } from 'src/service/user';

export default () => {
  const slide: Ref<string> = ref('style');

  const fetchUserList = async () => {
    try {
      const res = await getUserList();

      if (res.success) {
        console.log('res,', res.data);
      }
    } catch (error) {
      console.log(error);
    }
  };

  return {
    slide,
    fetchUserList,
  };
};
