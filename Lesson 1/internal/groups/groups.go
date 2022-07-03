package groups

import (
	"GO_back_2/Lesson 1/internal/users"
)

type groups struct {
	data []string
}

func (g *groups) AddGroup(Name string) (id int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (g *groups) AddUserInGroup(user users.User) error {
	//TODO implement me
	panic("implement me")
}

func (g *groups) FindUser(id int64) users.User {
	//TODO implement me
	panic("implement me")
}

func (g *groups) RemoveUser(id int64) {
	//TODO implement me
	panic("implement me")
}

func (g *groups) FindGroupByName(name string) {
	//TODO implement me
	panic("implement me")
}
