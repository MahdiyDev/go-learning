package services

import (
	"fmt"
	"tm/entities"
	"tm/lib"
)

type UserService struct {
	conn *lib.DatabaseConnection
}

func (u *UserService) New(conn *lib.DatabaseConnection) {
	u.conn = conn
}

func (u UserService) GetAllUsers() ([]entities.UserEntity, error) {
	userEntity := []entities.UserEntity{}
	err := u.conn.FetchAll(&userEntity, "select * from users")
	if err != nil {
		fmt.Printf("UserService GetAllUsers error: %v\n", err)
		return nil, err
	}
	return userEntity, nil
}

func (u UserService) GetOneUsers(id int) (*entities.UserEntity, error) {
	userEntity := entities.UserEntity{}
	err := u.conn.FetchOne(&userEntity, "select * from users where id = $1", id)
	if err != nil {
		fmt.Printf("UserService GetOneUsers error: %v\n", err)
		return nil, err
	}
	return &userEntity, nil
}

func (u UserService) CreateUser(userEntity entities.UserEntity) (*entities.UserEntity, error) {
	err := u.conn.FetchOne(&userEntity, "insert into users(name, age) values($1, $2) returning *", userEntity.Name, userEntity.Age)
	if err != nil {
		fmt.Printf("UserService GetAllUsers error: %v\n", err)
		return nil, err
	}
	return &userEntity, err
}
