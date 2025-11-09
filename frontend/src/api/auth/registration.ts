import { api } from '@/api';
import { IUser, IRegistration } from '@/interfaces';

interface RegistrationResult {
  user: IUser | null;
  access_token: string | null;
}

export const useRegistration = async (
  data: IRegistration,
): Promise<RegistrationResult> => {
  try {
    const response = await api.post('/auth/registration', data);

    return {
      user: response.data.user,
      access_token: response.data.accessToken,
    };
  } catch (error: unknown) {
    return {
      user: null,
      access_token: null,
    };
  }
};
