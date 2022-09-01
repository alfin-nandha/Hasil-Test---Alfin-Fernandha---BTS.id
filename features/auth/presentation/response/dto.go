package response

import (
	"btsid/test/features/auth"
)

type user struct {
	ID       int    `json:"id"`
	Username string `json:"email"`
	Password string `json:"password"`
}

func FromCore(data auth.Core) user {
	return user{
		ID:       data.ID,
		Username: data.Username,
		Password: data.Password,
	}
}
