package common

type BaseEntity struct {
	ID int
}

func (be *BaseEntity) SetID(id int) {
	be.ID = id
}
