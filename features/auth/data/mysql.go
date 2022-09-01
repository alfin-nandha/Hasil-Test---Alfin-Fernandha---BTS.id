package data

import (
	"btsid/test/features/auth"
	"errors"

	// "project/e-comerce/features/auth/data"

	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(conn *gorm.DB) auth.Data {
	return &mysqlAuthRepository{
		db: conn,
	}
}

func (repo *mysqlAuthRepository) FindUser(email string) (response auth.Core, err error) {
	dataUser := User{}
	result := repo.db.Where("username = ?", email).Find(&dataUser)
	if result.Error != nil || result.RowsAffected == 0 {
		return auth.Core{}, errors.New("user not found")
	}

	return dataUser.toCore(), nil
}

func (repo *mysqlAuthRepository) InsertUser(core auth.Core) error {
	authModel := fromCore(core)
	result := repo.db.Create(&authModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
