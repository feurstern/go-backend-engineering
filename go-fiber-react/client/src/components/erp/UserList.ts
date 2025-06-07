export default () => {
  const columns = [
    {
      name: 'no',
      label: 'no',
      //   align: 'left',
      field: 'no',
    },
    {
      name: 'username',
      label: 'username',
      //   align: 'left',
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

  return {
    columns,
    rows,
  };
};
