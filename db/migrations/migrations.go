package migrations

import (
	"fmt"

	"github.com/gemurdock/KeyFinder-GoLang/db/model"
	"gorm.io/gorm"
)

func MigrateAllModels(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return fmt.Errorf("failed to migrate models: %w", err)
	}

	return nil
}
