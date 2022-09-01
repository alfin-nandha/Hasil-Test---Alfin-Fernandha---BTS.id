package migration

import (
	auth "btsid/test/features/auth/data"
	checklist "btsid/test/features/checklist/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(auth.User{})
	db.AutoMigrate(checklist.Checklist{})
}
