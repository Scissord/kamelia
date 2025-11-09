import { api } from '@/api';
import { IUser, IUserLogin } from '@/interfaces';

interface LoginResult {
  user: IUser | null;
  accessToken: string | null;
}

export const useLogin = async (data: IUserLogin): Promise<LoginResult> => {
  try {
    const response = await api.post('/auth/login', data);

    return {
      user: response.data.user,
      accessToken: response.data.accessToken,
    };
  } catch (error: any) {
    return {
      user: null,
      accessToken: null,
    };
  }
};
