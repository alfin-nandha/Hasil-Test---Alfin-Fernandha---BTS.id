package business

import (
	"btsid/test/features/auth"
	"btsid/test/middlewares"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	userData auth.Data
}

func NewAuthBusiness(usrData auth.Data) auth.Business {
	return &authUseCase{
		userData: usrData,
	}
}

func (uc *authUseCase) Login(core auth.Core) (string, int, error) {
	response, errFind := uc.userData.FindUser(core.Username)
	if errFind != nil {
		return "", 0, errors.New("user not found")
	}
	errCompare := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(core.Password))
	if errCompare != nil {
		return "", 0, errors.New("wrong password")
	}
	token, err := middlewares.CreateToken(int(response.ID))

	return token, response.ID, err
}

func (uc *authUseCase) Register(core auth.Core) error {
	passWillBcrypt := []byte(core.Password)
	hash, err_hash := bcrypt.GenerateFromPassword(passWillBcrypt, bcrypt.DefaultCost)
	if err_hash != nil {
		return errors.New("hashing password failed")
	}
	core.Password = string(hash)
	response := uc.userData.InsertUser(core)
	if response != nil {
		return response
	}
	return nil
}
