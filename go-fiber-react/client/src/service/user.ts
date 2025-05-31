import { UserResponse } from 'src/components/models';
import api from 'src/utils/api';

export const getUserList = async (): Promise<UserResponse> => {
  const res = await api.get('/user');

  return res.data;
};
