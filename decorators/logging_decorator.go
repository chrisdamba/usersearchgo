package decorators

import (
	"github.com/chrisdamba/usersearchgo/db"
	"github.com/chrisdamba/usersearchgo/services"
	"log"
	"time"
)

type LoggingDecorator struct {
	next services.UserService
}

func (d *LoggingDecorator) SetNext(next services.UserService) {
	d.next = next
}

func (d *LoggingDecorator) GetUserByID(id int32) *db.User {
	start := time.Now()
	user := d.next.GetUserByID(id)
	log.Printf("GetUserByID(id=%d) took %v\n", id, time.Since(start))
	return user
}

func (d *LoggingDecorator) GetUsersByIDs(ids []int32) []*db.User {
	start := time.Now()
	users := d.next.GetUsersByIDs(ids)
	log.Printf("GetUsersByIDs(ids=%v) took %v\n", ids, time.Since(start))
	return users
}

func (d *LoggingDecorator) SearchUsers(fname, city string, phone string, married bool) []*db.User {
	start := time.Now()
	users := d.next.SearchUsers(fname, city, phone, married)
	log.Printf("SearchUsers(fname=%s, city=%s, phone=%d, married=%t) took %v\n", fname, city, phone, married, time.Since(start))
	return users
}
