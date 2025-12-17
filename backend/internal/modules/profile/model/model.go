package model

import "time"

type Profile struct {
	ID     int64 `gorm:"primaryKey;column:id" json:"id"`
	UserID int64 `gorm:"column:user_id;not null;index" json:"user_id"`

	FirstNameEncrypted []byte `gorm:"column:first_name_encrypted;type:bytea" json:"-"`
	FirstNameHash      string `gorm:"column:first_name_hash" json:"-"`

	LastNameEncrypted []byte `gorm:"column:last_name_encrypted;type:bytea" json:"-"`
	LastNameHash      string `gorm:"column:last_name_hash" json:"-"`

	MiddleNameEncrypted []byte `gorm:"column:middle_name_encrypted;type:bytea" json:"-"`
	MiddleNameHash      string `gorm:"column:middle_name_hash" json:"-"`

	EmailEncrypted []byte `gorm:"column:email_encrypted;type:bytea" json:"-"`
	EmailHash      string `gorm:"column:email_hash" json:"-"`

	PhoneEncrypted []byte `gorm:"column:phone_encrypted;type:bytea" json:"-"`
	PhoneHash      string `gorm:"column:phone_hash" json:"-"`

	AvatarURL string `gorm:"column:avatar_url" json:"avatar_url"`

	BirthdayEncrypted []byte `gorm:"column:birthday_encrypted;type:bytea" json:"-"`
	BirthdayHash      string `gorm:"column:birthday_hash" json:"-"`

	Gender   *string `gorm:"column:gender" json:"gender"`
	Locale   *string `gorm:"column:locale" json:"locale"`
	Timezone *string `gorm:"column:timezone" json:"timezone"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// Расшифрованные поля (НЕ сохраняются в БД)
	FirstName  *string `gorm:"-" json:"first_name,omitempty"`
	LastName   *string `gorm:"-" json:"last_name,omitempty"`
	MiddleName *string `gorm:"-" json:"middle_name,omitempty"`
	Email      *string `gorm:"-" json:"email,omitempty"`
	Phone      *string `gorm:"-" json:"phone,omitempty"`
	Birthday   *string `gorm:"-" json:"birthday,omitempty"`
}

func (Profile) TableName() string {
	return "auth.profile"
}
