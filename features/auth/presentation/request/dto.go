package request

import (
	"btsid/test/features/auth"
)

type User struct {
	Username string `json:"username" form:"username" validate:"required,username"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (u User) ToCore() auth.Core {
	userCore := auth.Core{
		Username: u.Username,
		Password: u.Password,
	}
	return userCore
}
