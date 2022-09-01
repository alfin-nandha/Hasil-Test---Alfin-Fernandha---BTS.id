package response

import (
	"btsid/test/features/checklist"
)

type Checklist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromCore(data checklist.Core) Checklist {
	return Checklist{
		ID:   data.ID,
		Name: data.Name,
	}
}
