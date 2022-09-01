package data

import (
	"btsid/test/features/checklist"

	"gorm.io/gorm"
)

type mysqlChecklistRepository struct {
	db *gorm.DB
}

func NewChecklistRepository(conn *gorm.DB) checklist.Data {
	return &mysqlChecklistRepository{
		db: conn,
	}
}

func (repo *mysqlChecklistRepository) FindData(id int) ([]checklist.Core, error) {
	checklistModel := []Checklist{}
	result := repo.db.Where("user_id = ?", id).Find(&checklistModel)
	if result.Error != nil {
		return []checklist.Core{}, result.Error
	}
	return toCoreList(checklistModel), nil
}
