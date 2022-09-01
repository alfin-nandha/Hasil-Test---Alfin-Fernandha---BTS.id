package request

import (
	"btsid/test/features/checklist"
)

type Checklist struct {
	UserID int
	Name   string `json:"name" form:"name" validate:"required"`
}

func (c Checklist) ToCore() checklist.Core {
	checklistCore := checklist.Core{
		UserID: c.UserID,
		Name:   c.Name,
	}
	return checklistCore
}
