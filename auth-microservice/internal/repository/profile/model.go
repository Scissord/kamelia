package profile

import "time"

type Profile struct {
	ID     int64 `db:"id" json:"id"`
	UserID int64 `db:"user_id" json:"user_id"`

	FirstNameEncrypted []byte `db:"first_name_encrypted" json:"-"`
	FirstNameHash      string `db:"first_name_hash" json:"-"`

	LastNameEncrypted []byte `db:"last_name_encrypted" json:"-"`
	LastNameHash      string `db:"last_name_hash" json:"-"`

	MiddleNameEncrypted []byte `db:"middle_name_encrypted" json:"-"`
	MiddleNameHash      string `db:"middle_name_hash" json:"-"`

	EmailEncrypted []byte `db:"email_encrypted" json:"-"`
	EmailHash      string `db:"email_hash" json:"-"`

	PhoneEncrypted []byte `db:"phone_encrypted" json:"-"`
	PhoneHash      string `db:"phone_hash" json:"-"`

	AvatarURL string `db:"avatar_url" json:"avatar_url"`

	BirthdayEncrypted []byte `db:"birthday_encrypted" json:"-"`
	BirthdayHash      string `db:"birthday_hash" json:"-"`

	Gender   string `db:"gender" json:"gender"` // ENUM: auth.gender_type
	Locale   string `db:"locale" json:"locale"` // ENUM: auth.locale_type
	Timezone string `db:"timezone" json:"timezone"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`

	FirstName  *string `db:"first_name" json:"first_name,omitempty"`
	LastName   *string `db:"last_name" json:"last_name,omitempty"`
	MiddleName *string `db:"middle_name" json:"middle_name,omitempty"`
	Email      *string `db:"email" json:"email,omitempty"`
	Phone      *string `db:"phone" json:"phone,omitempty"`
	Birthday   *string `db:"birthday" json:"birthday,omitempty"`
}

type ProfileCreateInput struct {
	UserID     int64
	FirstName  *string
	LastName   *string
	MiddleName *string
	Email      *string
	Phone      *string
	Birthday   *string
	Gender     *string
	Locale     *string
	Timezone   *string
}
