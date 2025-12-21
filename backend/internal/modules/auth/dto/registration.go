package dto

type RegistrationUserDTO struct {
	Login      string  `json:"login" validate:"required,min=3,max=32"`
	Password   string  `json:"password" validate:"required,min=8,max=128"`
	FirstName  *string `json:"first_name,omitempty" validate:"omitempty,min=2,max=255"`
	LastName   *string `json:"last_name,omitempty" validate:"omitempty,min=2,max=255"`
	MiddleName *string `json:"middle_name,omitempty" validate:"omitempty,min=2,max=255"`
	Email      *string `json:"email,omitempty" validate:"omitempty,email,max=255"`
	Phone      *string `json:"phone,omitempty" validate:"omitempty,e164"`
	Birthday   *string `json:"birthday,omitempty" validate:"omitempty,len=10,datetime=2006-01-02"`
	Gender     *string `json:"gender" validate:"required,oneof=male female other"`
	Locale     *string `json:"locale" validate:"required"`
	Timezone   *string `json:"timezone" validate:"required,max=255"`
	AvatarURL  *string `json:"avatar_url,omitempty" validate:"omitempty,max=255"`
}
