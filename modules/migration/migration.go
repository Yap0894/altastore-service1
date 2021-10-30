package migration

import (
	"AltaStore/modules/admin"
	"AltaStore/modules/user"

	"gorm.io/gorm"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&admin.Admin{},
	)
}
