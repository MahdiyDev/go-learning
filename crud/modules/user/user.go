package user

import (
	"crud/common"
	"fmt"
)

type UserEntity struct {
	Name string
	Age  int

	common.BaseEntity
}

var Users []UserEntity

func (u *UserEntity) New() {
	uLen := len(Users)
	if uLen == 0 {
		u.SetID(1)
	} else {
		u.SetID(Users[len(Users)-1].ID + 1)
	}
}

func (u *UserEntity) SetName(name string) {
	u.Name = name
}

func (u *UserEntity) SetAge(age int) {
	u.Age = age
}

func (u *UserEntity) Append() {
	Users = append(Users, *u)
}


func (u UserEntity) String() string {
	return fmt.Sprintf("ID: %v | Name: %v | Age: %v |", u.ID, u.Name, u.Age)
}
