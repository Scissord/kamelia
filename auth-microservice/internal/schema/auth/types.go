package auth

import "time"

type RegistrationInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`

	FirstName  *string `json:"first_name,omitempty"`
	LastName   *string `json:"last_name,omitempty"`
	MiddleName *string `json:"middle_name,omitempty"`
	Email      *string `json:"email,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	Birthday   *string `json:"birthday,omitempty"`
	Gender     *string `json:"gender,omitempty"`
	Locale     *string `json:"locale,omitempty"`
	Timezone   *string `json:"timezone,omitempty"`
}

type RegistrationResponse struct {
	ID         int64      `json:"id"`
	Login      string     `json:"login"`
	FirstName  *string    `json:"first_name,omitempty"`
	MiddleName *string    `json:"middle_name,omitempty"`
	LastName   *string    `json:"last_name,omitempty"`
	Email      *string    `json:"email,omitempty"`
	Phone      *string    `json:"phone,omitempty"`
	Birthday   *string    `json:"birthday,omitempty"`
	Gender     *string    `json:"gender,omitempty"`
	Locale     *string    `json:"locale,omitempty"`
	Timezone   *string    `json:"timezone,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}
