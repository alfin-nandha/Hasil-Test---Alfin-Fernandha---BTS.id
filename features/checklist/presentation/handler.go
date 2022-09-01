package presentation

import (
	"btsid/test/features/checklist"
	"btsid/test/helper"
	"btsid/test/middlewares"

	"github.com/labstack/echo/v4"
)

type ChecklistHandler struct {
	checklistBusiness checklist.Business
}

func NewChecklistHandler(business checklist.Business) *ChecklistHandler {
	return &ChecklistHandler{
		checklistBusiness: business,
	}
}

func (h *ChecklistHandler) GetData(c echo.Context) error {
	userID, errToken := middlewares.ExtractToken(c)
	if userID == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("not authorized"))
	}

	result, err := h.checklistBusiness.GetData(userID)

	if err != nil {
		return c.JSON(helper.ResponseNotFound(err.Error()))
	}
	return c.JSON(helper.ResponseStatusOkWithData("get data success", result))
}

// func (h *AuthHandler) Register(c echo.Context) error {
// 	reqBody := request.User{}
// 	errBind := c.Bind(&reqBody)
// 	if errBind != nil {
// 		return c.JSON(helper.ResponseBadRequest("bad request"))
// 	}
// 	err := h.userBusiness.Register(reqBody.ToCore())
// 	if err != nil {
// 		return c.JSON(helper.ResponseBadRequest("bad request"))
// 	}
// 	return c.JSON(helper.ResponseCreateSuccess("register success"))
// }
