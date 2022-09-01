package checklist

type Core struct {
	ID     int
	UserID int
	Name   string
}

type Data interface {
	FindData(id int) ([]Core, error)
}

type Business interface {
	GetData(id int) ([]Core, error)
}
