package utils

import "gorm.io/gorm"

func Encrypt(value *string, key string) interface{} {
	if value == nil {
		return nil
	}
	return gorm.Expr("pgp_sym_encrypt(?, ?)", *value, key)
}
