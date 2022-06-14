package communities

import "GO_back_2/Lesson 1/internal/users"

type comm struct {
	CommName string
}

func (c *comm) AddGroup(Name string) (id int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *comm) AddUserInGroup(user users.User) error {
	//TODO implement me
	panic("implement me")
}

func (c *comm) FindUser(id int64) users.User {
	//TODO implement me
	panic("implement me")
}

func (c *comm) RemoveUser(id int64) {
	//TODO implement me
	panic("implement me")
}

func (c *comm) FindGroupByName(name string) {
	//TODO implement me
	panic("implement me")
}
