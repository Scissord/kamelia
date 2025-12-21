import { IUser, IUserLogin, IRegistration } from '@/interfaces';
import { base_url } from '@/utils';

interface LoginResult {
  user: IUser | null;
  access_token: string | null;
}

type ApiError = {
  error: string;
};

export const AuthService = {
  async registration(data: IRegistration): Promise<IUser | string> {
    try {
      const response = await fetch(`${base_url}/auth/registration`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });

      const result = await response.json();
      if (!response.ok) {
        throw result;
      }

      return result;
    } catch (err: unknown) {
      const apiErr = err as ApiError;
      if (apiErr?.error) {
        return apiErr.error; // "EMAIL_EXISTS" | "PHONE_EXISTS"
      }

      return 'UNKNOWN_ERROR';
    }
  },

  async login(data: IUserLogin): Promise<LoginResult> {
    try {
      const response = await fetch(`${base_url}/auth/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });

      if (!response.ok) {
        throw new Error('Login failed');
      }

      const result = await response.json();

      return {
        user: result.user,
        access_token: result.accessToken,
      };
    } catch (error: unknown) {
      return {
        user: null,
        access_token: null,
      };
    }
  },

  async logout() {},

  async refresh() {},
};
