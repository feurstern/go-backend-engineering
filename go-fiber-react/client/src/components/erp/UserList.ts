import { getUserList } from 'src/service/user';
import { UserListRows } from '../models';
import { ref } from 'vue';

export default () => {
  const columns = [
    {
      name: 'no',
      label: 'no',
      // align: 'left',
      field: 'no',
    },
    {
      name: 'username',
      label: 'username',
      // align: 'center',
      field: 'Username',
    },
    {
      name: 'email',
      label: 'email',
      //   align: 'left',
      field: 'Email',
    },
  ];

  const rows = [
    {
      no: 1,
      Username: 'bbi',
      Email: 'mimkui@.ds',
    },
    {
      no: 2,
      Username: 'babi',
      Email: 'babi@.ds',
    },
  ];

  const rows2 = ref<UserListRows[]>([]);

  const fetchUserList = async () => {
    try {
      const res = await getUserList();
      if (res.success) {
        rows2.value = res.data.map((x, index) => ({
          no: index + 1,
          Username: x.Username,
          Email: x.Email,
        }));
      }
    } catch (error) {
      alert('errori');
    }
  };

  console.log('res', rows2.value.length);

  return {
    columns,
    rows,
    rows2,
    fetchUserList,
  };
};
