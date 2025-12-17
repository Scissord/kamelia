package migrations

import (
	client "backend/internal/modules/client/model"

	"gorm.io/gorm"
)

func AutoMigrateAll(db *gorm.DB) error {
	// client
	if err := db.AutoMigrate(&client.Client{}); err != nil {
		return err
	}

	return nil
}
