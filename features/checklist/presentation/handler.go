package presentation

import (
	"btsid/test/features/checklist"
	"btsid/test/features/checklist/presentation/request"
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

func (h *ChecklistHandler) InsertData(c echo.Context) error {
	userID, errToken := middlewares.ExtractToken(c)
	if userID == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("not authorized"))
	}
	reqBody := request.Checklist{}
	errBind := c.Bind(&reqBody)
	if errBind != nil {
		return c.JSON(helper.ResponseBadRequest("bad request"))
	}
	reqBody.UserID = userID
	err := h.checklistBusiness.InsertData(reqBody.ToCore())
	if err != nil {
		return c.JSON(helper.ResponseBadRequest("bad request"))
	}
	return c.JSON(helper.ResponseCreateSuccess("insert success"))
}
