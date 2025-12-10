package migrations

import (
	client "backend/internal/client/model"

	"gorm.io/gorm"
)

func AutoMigrateAll(db *gorm.DB) error {
	// client
	if err := db.AutoMigrate(&client.Client{}); err != nil {
		return err
	}

	return nil
}
