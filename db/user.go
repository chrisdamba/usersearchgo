package db

import (
	"github.com/jaswdr/faker/v2"
	"sync"
)

type User struct {
	Id      int32   `faker:"unique_id"`
	Fname   string  `faker:"first_name"`
	City    string  `faker:"city"`
	Phone   string  `faker:"phone_number"`
	Height  float32 `faker:"float32,between:5.0,7.0"`
	Married bool    `faker:"bool"`
}

type UserStore struct {
	users []User
}

var instance *UserStore
var once sync.Once

func GetInstance() *UserStore {
	once.Do(func() {
		instance = &UserStore{}
		instance.generateFakeUsers(100) // Generate 100 fake users
	})
	return instance
}

func (store *UserStore) generateFakeUsers(num int) {
	fake := faker.New()
	for i := 0; i < num; i++ {
		user := User{
			Id:      int32(i + 1),
			Fname:   fake.Person().FirstName(),
			City:    fake.Address().City(),
			Phone:   fake.Phone().Number(),
			Height:  fake.Float32(2, 4, 7),
			Married: fake.Bool(),
		}
		store.users = append(store.users, user)
	}
}

func (store *UserStore) GetUserByID(id int32) *User {
	for _, user := range store.users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func (store *UserStore) GetUsersByIDs(ids []int32) []*User {
	var foundUsers []*User
	for _, id := range ids {
		user := store.GetUserByID(id)
		if user != nil {
			foundUsers = append(foundUsers, user)
		}
	}
	return foundUsers
}

func (store *UserStore) SearchUsers(fname, city string, phone string, married bool) []*User {
	var foundUsers []*User
	for _, user := range store.users {
		if (fname == "" || user.Fname == fname) &&
			(city == "" || user.City == city) &&
			(phone == "" || user.Phone == phone) &&
			(!married || user.Married == married) {
			foundUsers = append(foundUsers, &user)
		}
	}
	return foundUsers
}
