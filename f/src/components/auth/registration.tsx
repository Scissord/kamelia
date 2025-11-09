'use client';

import { useState } from 'react';
import { useForm, Controller } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import {
  Label,
  Input,
  Button,
  Popover,
  PopoverTrigger,
  PopoverContent,
  Calendar,
  RadioGroup,
  RadioGroupItem,
} from '@/components';
import { useRegistration } from '@/api';
import { ChevronDownIcon, EyeIcon, EyeOffIcon } from 'lucide-react';

const registrationSchema = z.object({
  login: z.string().min(3, 'Логин должен быть не менее 3 символов'),
  password: z.string().min(6, 'Пароль должен быть не менее 6 символов'),
  email: z.string().email('Некорректный email').optional().or(z.literal('')),
  phone: z
    .string()
    .regex(/^\+?\d{10,15}$/, 'Некорректный телефон')
    .optional()
    .or(z.literal('')),
  birthday: z.date().optional(),
  gender: z.enum(['male', 'female', 'other']).optional(),
});

type RegistrationFormData = z.infer<typeof registrationSchema>;

export const RegistrationForm = () => {
  const [open, setOpen] = useState(false);
  const [showPassword, setShowPassword] = useState(false);

  const {
    register,
    handleSubmit,
    control,
    formState: { errors, isSubmitting },
  } = useForm<RegistrationFormData>({
    resolver: zodResolver(registrationSchema),
    mode: 'onBlur',
  });

  const onSubmit = async (data: RegistrationFormData) => {
    try {
      const locale = navigator.language || navigator.languages[0] || 'ru';
      const timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;

      const user = await useRegistration({
        ...data,
        locale,
        timezone,
      });

      console.log('Успешная регистрация', user);
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <form
      className="space-y-4 p-4"
      onSubmit={handleSubmit(onSubmit)}
      noValidate
    >
      <div>
        <Label htmlFor="login">Логин</Label>
        <Input
          id="login"
          placeholder="tester"
          {...register('login')}
          className={errors.login ? 'border-red-500' : ''}
        />
        {errors.login && (
          <p className="text-red-500 text-sm mt-1">{errors.login.message}</p>
        )}
      </div>

      <div>
        <Label htmlFor="password">Пароль</Label>
        <div className="relative">
          <Input
            id="password"
            placeholder="••••••••"
            type={showPassword ? 'text' : 'password'}
            {...register('password')}
            className={errors.password ? 'border-red-500' : ''}
          />
          <button
            // variant={'ghost'}
            // size={'sm'}
            type="button"
            className="absolute right-2 top-1/2 -translate-y-1/2 text-gray-500"
            onClick={() => setShowPassword(!showPassword)}
          >
            {showPassword ? <EyeOffIcon size={16} /> : <EyeIcon size={16} />}
          </button>
        </div>
        {errors.password && (
          <p className="text-red-500 text-sm mt-1">{errors.password.message}</p>
        )}
      </div>

      <div>
        <Label htmlFor="email">Email</Label>
        <Input
          id="email"
          type="email"
          placeholder="example@mail.com"
          {...register('email')}
          className={errors.email ? 'border-red-500' : ''}
        />
        {errors.email && (
          <p className="text-red-500 text-sm mt-1">{errors.email.message}</p>
        )}
      </div>

      <div>
        <Label htmlFor="phone">Телефон</Label>
        <Input
          id="phone"
          placeholder="+77777777777"
          {...register('phone')}
          className={errors.phone ? 'border-red-500' : ''}
        />
        {errors.phone && (
          <p className="text-red-500 text-sm mt-1">{errors.phone.message}</p>
        )}
      </div>

      <div className="flex flex-col gap-1">
        <Label htmlFor="birthday">Дата рождения</Label>
        <Controller
          control={control}
          name="birthday"
          render={({ field }) => {
            const date = field.value;
            return (
              <>
                <Popover open={open} onOpenChange={setOpen}>
                  <PopoverTrigger asChild>
                    <Button
                      variant="outline"
                      id="birthday"
                      className="w-48 justify-between font-normal"
                    >
                      {date ? date.toLocaleDateString() : 'Выберите дату'}
                      <ChevronDownIcon />
                    </Button>
                  </PopoverTrigger>
                  <PopoverContent className="w-64 p-2" align="start">
                    <Calendar
                      mode="single"
                      selected={date}
                      onSelect={(selectedDate) => {
                        field.onChange(selectedDate);
                        setOpen(false);
                      }}
                      captionLayout="dropdown"
                      className="w-full"
                    />
                  </PopoverContent>
                </Popover>
                {errors.birthday && (
                  <p className="text-red-500 text-sm mt-1">
                    {errors.birthday.message}
                  </p>
                )}
              </>
            );
          }}
        />
      </div>

      <div>
        <Label htmlFor="gender">Пол</Label>
        <Controller
          control={control}
          name="gender"
          render={({ field }) => (
            <RadioGroup
              onValueChange={field.onChange}
              value={field.value}
              className="flex gap-4 mt-1"
            >
              <div className="flex items-center gap-1">
                <RadioGroupItem value="male" id="gender-male" />
                <Label htmlFor="gender-male">Мужской</Label>
              </div>
              <div className="flex items-center gap-1">
                <RadioGroupItem value="female" id="gender-female" />
                <Label htmlFor="gender-female">Женский</Label>
              </div>
              <div className="flex items-center gap-1">
                <RadioGroupItem value="other" id="gender-other" />
                <Label htmlFor="gender-other">Другое</Label>
              </div>
            </RadioGroup>
          )}
        />
        {errors.gender && (
          <p className="text-red-500 text-sm mt-1">{errors.gender.message}</p>
        )}
      </div>

      <div className="flex items-center justify-between">
        <Button variant="link" className="ml-auto p-0">
          Уже есть аккаунт?
        </Button>
      </div>

      <Button className="w-full mt-2" type="submit" disabled={isSubmitting}>
        {isSubmitting ? 'Загрузка...' : 'Зарегистрироваться'}
      </Button>
    </form>
  );
};
