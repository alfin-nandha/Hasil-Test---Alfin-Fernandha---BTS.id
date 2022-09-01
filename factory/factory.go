package factory

import (
	checklistBusiness "btsid/test/features/checklist/bussiness"
	checklistData "btsid/test/features/checklist/data"
	checklistPresentation "btsid/test/features/checklist/presentation"

	authBusiness "btsid/test/features/auth/business"
	authData "btsid/test/features/auth/data"
	authPresentation "btsid/test/features/auth/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	ChecklistPresenter *checklistPresentation.ChecklistHandler
	AuthPresenter      *authPresentation.AuthHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {

	checklistData := checklistData.NewChecklistRepository(dbConn)
	checklistBusiness := checklistBusiness.NewProductBusiness(checklistData)
	checklistPresentation := checklistPresentation.NewChecklistHandler(checklistBusiness)

	authData := authData.NewAuthRepository(dbConn)
	authBusiness := authBusiness.NewAuthBusiness(authData)
	authPresentation := authPresentation.NewAuthHandler(authBusiness)

	return Presenter{
		ChecklistPresenter: checklistPresentation,
		AuthPresenter:      authPresentation,
	}
}
