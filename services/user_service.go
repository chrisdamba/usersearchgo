package services

import (
	"github.com/chrisdamba/usersearchgo/db"
	"github.com/chrisdamba/usersearchgo/factories"
)

type UserService struct {
	store   *db.UserStore
	factory *factories.UserFactory
}

func NewUserService() *UserService {
	return &UserService{
		store:   db.GetInstance(),
		factory: &factories.UserFactory{},
	}
}

func (us *UserService) GetUserByID(id int32) *db.User {
	return us.store.GetUserByID(id)
}

func (us *UserService) GetUsersByIDs(ids []int32) []*db.User {
	return us.store.GetUsersByIDs(ids)
}

func (us *UserService) SearchUsers(fname, city string, phone string, married bool) []*db.User {
	return us.store.SearchUsers(fname, city, phone, married)
}
