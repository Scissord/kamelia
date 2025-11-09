'use client';

import { useState } from 'react';
import {
  Card,
  CardHeader,
  CardContent,
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
  LoginForm,
  RegistrationForm,
} from '@/components';

export function AuthenticationForm() {
  const [tab, setTab] = useState('login');

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <Card className="w-[400px]">
        <Tabs value={tab} onValueChange={setTab} className="w-full">
          <CardHeader className="pb-0">
            <TabsList>
              <TabsTrigger value="login">Вход</TabsTrigger>
              <TabsTrigger value="register">Регистрация</TabsTrigger>
            </TabsList>
          </CardHeader>

          <CardContent className="pt-4">
            <TabsContent value="login">
              <LoginForm />
            </TabsContent>
            <TabsContent value="register">
              <RegistrationForm />
            </TabsContent>
          </CardContent>
        </Tabs>
      </Card>
    </div>
  );
}
