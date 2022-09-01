package presentation

import (
	"btsid/test/features/auth"
	"btsid/test/features/auth/presentation/request"
	"btsid/test/helper"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userBusiness auth.Business
}

func NewAuthHandler(business auth.Business) *AuthHandler {
	return &AuthHandler{
		userBusiness: business,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	reqBody := request.User{}
	errBind := c.Bind(&reqBody)
	if errBind != nil {
		return c.JSON(helper.ResponseBadRequest("bad request"))
	}

	authCore := reqBody.ToCore()
	token, id, err := h.userBusiness.Login(authCore)

	if err != nil {
		return c.JSON(helper.ResponseNotFound(err.Error()))
	}
	return c.JSON(helper.ResponseStatusOkWithData("login success", map[string]interface{}{
		"id":    id,
		"token": token,
	}))
}

func (h *AuthHandler) Register(c echo.Context) error {
	reqBody := request.User{}
	errBind := c.Bind(&reqBody)
	if errBind != nil {
		return c.JSON(helper.ResponseBadRequest("bad request"))
	}
	err := h.userBusiness.Register(reqBody.ToCore())
	if err != nil {
		return c.JSON(helper.ResponseBadRequest("bad request"))
	}
	return c.JSON(helper.ResponseCreateSuccess("register success"))
}
