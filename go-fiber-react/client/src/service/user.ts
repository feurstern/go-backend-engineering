import {
  LoginPayload,
  LoginResponse,
  userRegistrationPayload,
  UserRegistrationRespsonse,
  UserResponse,
} from 'src/components/models';
import api from 'src/utils/api';

export const getUserList = async (): Promise<UserResponse> => {
  const res = await api.get('/user');
  return res.data;
};

export const userRegistration = async (
  payload: userRegistrationPayload,
): Promise<UserRegistrationRespsonse> => {
  const res = await api.post('/registration', payload);

  return res.data;
};

export const login = async (payload: LoginPayload): Promise<LoginResponse> => {
  const res = await api.post('/login', payload, {});
  return res.data;
};
