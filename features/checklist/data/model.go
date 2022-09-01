package data

import (
	"btsid/test/features/checklist"

	"gorm.io/gorm"
)

type Checklist struct {
	gorm.Model
	UserID int
	Name   string
	User   User
}

type User struct {
	gorm.Model
	Username string
	Password string
}

//DTO

func (data *Checklist) toCore() checklist.Core {
	return checklist.Core{
		ID:     int(data.ID),
		UserID: data.UserID,
		Name:   data.Name,
	}
}

func toCoreList(data []Checklist) []checklist.Core {
	result := []checklist.Core{}
	for ind := range data {
		result = append(result, data[ind].toCore())
	}
	return result
}

func fromCore(core checklist.Core) Checklist {
	return Checklist{
		UserID: core.ID,
		Name:   core.Name,
	}
}
