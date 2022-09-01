package data

import (
	"btsid/test/features/auth"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

//DTO

func (data *User) toCore() auth.Core {
	return auth.Core{
		ID:        int(data.ID),
		Username:  data.Username,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func fromCore(core auth.Core) User {
	return User{
		Username: core.Username,
		Password: core.Password,
	}
}
