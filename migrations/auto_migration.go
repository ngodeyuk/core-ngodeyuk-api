package migrations

import (
	"ngodeyuk-core/pkg/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Course{},
	)
	if err != nil {
		return err
	}
	return nil
}
