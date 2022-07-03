package users

import "os/user"

type User struct {
	_    struct{}
	Name string
}

type Users struct {
	list []User
}

func NewUsers() *Users {
	return &Users{}
}

func (u Users) AddUser(user User) (id int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (u Users) EditUser(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (u Users) RemoveUser(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (u Users) GetUser(id int64) user.User {
	//TODO implement me
	panic("implement me")
}
