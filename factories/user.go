package factories

import (
	"github.com/chrisdamba/usersearchgo/db"
	"github.com/jaswdr/faker/v2"
)

type UserFactory struct{}

func (uf *UserFactory) CreateUser(id int32) db.User {
	fake := faker.New()
	user := db.User{
		Id:      id,
		Fname:   fake.Person().FirstName(),
		City:    fake.Address().City(),
		Phone:   fake.Phone().Number(),
		Height:  fake.Float32(2, 4, 7),
		Married: fake.Bool(),
	}
	return user
}
