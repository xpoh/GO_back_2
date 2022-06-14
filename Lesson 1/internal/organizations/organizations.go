package organizations

import "GO_back_2/Lesson 1/internal/users"

type Organization struct {
	users []users.User
}

type orgs struct {
	list []Organization
}

func (o orgs) AddGroup(Name string) (id int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (o orgs) AddUserInGroup(user users.User) error {
	//TODO implement me
	panic("implement me")
}

func (o orgs) FindUser(id int64) users.User {
	//TODO implement me
	panic("implement me")
}

func (o orgs) RemoveUser(id int64) {
	//TODO implement me
	panic("implement me")
}

func (o orgs) FindGroupByName(name string) {
	//TODO implement me
	panic("implement me")
}

func NewOrgs() *orgs {
	return &orgs{}
}
