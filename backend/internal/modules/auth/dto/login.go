package dto

// Звёздочка нужна только тогда, когда тебе важно различать:

// «поле не пришло» vs «поле пришло, но пустое»

// Для логина это не нужно.

// На твоём примере
// Сейчас у тебя
// type LoginUserDTO struct {
// 	Login    *string
// 	Password string
// }

// Это означает:

// Ситуация в JSON	Login
// "login": "admin"	указатель → "admin"
// "login": ""	указатель → ""
// "login": null	nil
// поля вообще нет	nil

// Но для логина тебе не важно ни null, ни отсутствие поля —
// это всегда ошибка.

type LoginUserDTO struct {
	Login    *string `json:"login" validate:"required,min=3,max=32"`
	Password string  `json:"password" validate:"required,min=8,max=128"`
}
