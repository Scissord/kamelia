import { IUser, IUserLogin, IRegistration } from '@/interfaces';

interface LoginResult {
  user: IUser | null;
  access_token: string | null;
}

interface IRegistrationResult {
  user: IUser | null;
  access_token: string | null;
}

export const AuthService = {
  async registration(data: IRegistration): Promise<IRegistrationResult> {
    try {
      // Преобразуем данные для отправки на бэкенд
      const payload: any = {
        login: data.login,
        password: data.password,
      };

      // Добавляем опциональные поля только если они заполнены
      if (data.email && data.email.trim()) {
        payload.email = data.email.trim();
      }
      if (data.phone && data.phone.trim()) {
        payload.phone = data.phone.trim();
      }
      if (data.birthday instanceof Date) {
        payload.birthday = data.birthday.toISOString().split('T')[0];
      }
      if (data.gender) {
        payload.gender = data.gender;
      }
      if (data.locale) {
        payload.locale = data.locale;
      }
      if (data.timezone) {
        payload.timezone = data.timezone;
      }

      console.log('Отправка payload на бэкенд:', payload);

      const response = await fetch(
        'http://localhost:8080/api/auth/registration',
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(payload),
        },
      );

      if (!response.ok) {
        const errorText = await response.text();
        console.error('Registration failed:', response.status, errorText);
        throw new Error(`Registration failed: ${errorText}`);
      }

      const result = await response.json();
      console.log('Registration response:', result);

      // Бэкенд возвращает напрямую user объект, а не {user, accessToken}
      // Проверяем, что это действительно user объект
      if (result && result.id) {
        return {
          user: result,
          access_token: null, // Бэкенд не возвращает токен при регистрации
        };
      }

      // Если формат неожиданный, пробуем найти user в ответе
      return {
        user: result.user || result,
        access_token: result.accessToken || null,
      };
    } catch (error: unknown) {
      console.error('Registration error:', error);
      return {
        user: null,
        access_token: null,
      };
    }
  },

  async login(data: IUserLogin): Promise<LoginResult> {
    try {
      const response = await fetch('http://localhost:8080/api/auth/login', {
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
