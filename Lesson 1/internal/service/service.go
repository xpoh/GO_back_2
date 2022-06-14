package service

import (
	"GO_back_2/Lesson 1/internal/users"
	"os/user"
)

type IUsers interface {
	AddUser(user users.User) (id int64, err error)
	EditUser(id int64) error
	RemoveUser(id int64) error
	GetUser(id int64) user.User
}

type Grouper interface {
	AddGroup(Name string) (id int64, err error)
	AddUserInGroup(user users.User) error
	FindUser(id int64) users.User
	RemoveUser(id int64)
	FindGroupByName(name string) ()
}

type Service struct {
	users     IUsers
	groups    Grouper
	community Grouper
	orgs      Grouper
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) AddUser(user users.User) (id int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) AddGroup(name string) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RemoveUserFromGroup(name string) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) FindUser(name string) user.User {
	//TODO implement me
	panic("implement me")
}

func (s *Service) FindGroupWithUser(name string) string {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Init() {

}
