package bussiness

import (
	"btsid/test/features/checklist"
)

type checklistUseCase struct {
	checklistData checklist.Data
}

func NewProductBusiness(data checklist.Data) checklist.Business {
	return &checklistUseCase{
		checklistData: data,
	}
}

func (uc *checklistUseCase) GetData(id int) ([]checklist.Core, error) {
	result, err := uc.checklistData.FindData(id)
	if err != nil {
		return []checklist.Core{}, err
	}
	return result, nil
}

func (uc *checklistUseCase) InsertData(core checklist.Core) error {
	err := uc.checklistData.InsertData(core)
	if err != nil {
		return err
	}
	return nil
}
