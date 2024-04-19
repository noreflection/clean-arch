package migrations

import (
	"go-cqrs/internal/domain"
	"gorm.io/gorm"
)

type InitialCreateMigration struct{}

func (m *InitialCreateMigration) MigrationID() string {
	return "CreateUserTable"
}

func (m *InitialCreateMigration) Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&domain.Order{})
	if err != nil {
		return err
	}
	return nil
}

func (m *InitialCreateMigration) Rollback(db *gorm.DB) error {
	err := db.Migrator().DropTable(&domain.Order{})
	if err != nil {
		return err
	}
	return nil
}
