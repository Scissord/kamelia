package dto

type RegistrationUserDTO struct {
	Login    string `json:"login" validate:"required,min=2"`
	Password string `json:"password" validate:"required,min=2"`

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
