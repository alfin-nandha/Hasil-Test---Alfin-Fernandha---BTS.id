package migration

import (
	auth "btsid/test/features/auth/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(auth.User{})
}
