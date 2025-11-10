package profile

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	utils "auth-microservice/internal/utils"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Create(profile *ProfileCreateInput) (*Profile, error) {
	gender := "other"
	if profile.Gender != nil {
		gender = *profile.Gender
	}

	locale := "ru"
	if profile.Locale != nil {
		locale = *profile.Locale
	}

	query := `
		INSERT INTO auth."profile" (
			user_id,
			first_name_encrypted,
			first_name_hash,
			last_name_encrypted,
			last_name_hash,
			middle_name_encrypted,
			middle_name_hash,
			email_encrypted,
			email_hash,
			phone_encrypted,
			phone_hash,
			birthday_encrypted,
			birthday_hash,
			gender,
			locale,
			timezone
		)
		VALUES (
			$2,
			pgp_sym_encrypt(COALESCE(NULLIF(trim(COALESCE($3, '')), ''), ''), $1),
			CASE WHEN COALESCE($3, '') != '' THEN encode(digest(upper(trim(COALESCE($3, ''))), 'sha256'), 'hex') ELSE NULL END
			pgp_sym_encrypt(COALESCE(NULLIF(trim(COALESCE($4, '')), ''), ''), $1),
			CASE WHEN COALESCE($4, '') != '' THEN encode(digest(upper(trim(COALESCE($4, ''))), 'sha256'), 'hex') ELSE NULL END
			pgp_sym_encrypt(COALESCE(NULLIF(trim(COALESCE($5, '')), ''), ''), $1),
			CASE WHEN COALESCE($5, '') != '' THEN encode(digest(upper(trim(COALESCE($5, ''))), 'sha256'), 'hex') ELSE NULL END
			pgp_sym_encrypt(COALESCE(NULLIF(trim(COALESCE($6, '')), ''), ''), $1),
			CASE WHEN COALESCE($6, '') != '' THEN encode(digest(upper(trim(COALESCE($6, ''))), 'sha256'), 'hex') ELSE NULL END
			pgp_sym_encrypt(COALESCE(NULLIF(trim(COALESCE($7, '')), ''), ''), $1),
			CASE WHEN COALESCE($7, '') != '' THEN encode(digest(upper(trim(COALESCE($7, ''))), 'sha256'), 'hex') ELSE NULL END
			pgp_sym_encrypt(COALESCE(NULLIF(trim(COALESCE($8, '')), ''), ''), $1),
			CASE WHEN COALESCE($8, '') != '' THEN encode(digest(upper(trim(COALESCE($8, ''))), 'sha256'), 'hex') ELSE NULL END
			$9,
			$10,
			$11
		)
		RETURNING
			id,
			user_id,
			pgp_sym_decrypt(first_name_encrypted, $1) AS first_name,
			pgp_sym_decrypt(last_name_encrypted, $1) AS last_name,
			pgp_sym_decrypt(middle_name_encrypted, $1) AS middle_name,
			pgp_sym_decrypt(email_encrypted, $1) AS email,
			pgp_sym_decrypt(phone_encrypted, $1) AS phone,
			pgp_sym_decrypt(birthday_encrypted, $1) AS birthday,
			gender,
			locale,
			timezone,
			created_at,
			updated_at
	`

	secretKey := os.Getenv("SECRET_KEY")
	var createdProfile Profile
	profileJSON, _ := json.Marshal(profile)
	log.Printf("Profile before created: %s", profileJSON)

	// 2025/11/10 21:31:44 Profile before created: {"UserID":14,"FirstName":null,"LastName":null,"MiddleName":null,"Email":"test@email.com","Phone":null,"Birthday":null,"Gender":null,"Locale":null,"Timezone":null}

	err := r.DB.QueryRow(
		query,
		secretKey,
		profile.UserID,
		utils.StringOrEmpty(profile.FirstName),
		utils.StringOrEmpty(profile.LastName),
		utils.StringOrEmpty(profile.MiddleName),
		utils.StringOrEmpty(profile.Email),
		utils.StringOrEmpty(profile.Phone),
		utils.StringOrEmpty(profile.Birthday),
		gender,
		locale,
		utils.StringOrEmpty(profile.Timezone),
	).Scan(
		&createdProfile.ID,
		&createdProfile.UserID,
		&createdProfile.FirstName,
		&createdProfile.LastName,
		&createdProfile.MiddleName,
		&createdProfile.Email,
		&createdProfile.Phone,
		&createdProfile.Birthday,
		&createdProfile.Gender,
		&createdProfile.Locale,
		&createdProfile.Timezone,
		&createdProfile.CreatedAt,
		&createdProfile.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &createdProfile, nil
}

// field_encrypted = pgp_sym_encrypt($13, '${process.env.SECRET_KEY}'),
// field_code_hash = encode(digest(upper($13), 'sha256'), 'hex'),
