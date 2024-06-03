package decorators

import (
	"github.com/chrisdamba/usersearchgo/db"
	"github.com/chrisdamba/usersearchgo/services"
	"sync"
)

type CachingDecorator struct {
	next  services.UserService
	Cache map[int32]*db.User
	mu    sync.Mutex
}

// SetNext sets the next decorator or the base service.
func (d *CachingDecorator) SetNext(next services.UserService) {
	d.next = next
}

func (d *CachingDecorator) GetUserByID(id int32) *db.User {
	d.mu.Lock()
	defer d.mu.Unlock()
	if user, exists := d.Cache[id]; exists {
		return user
	}
	user := d.next.GetUserByID(id)
	if user != nil {
		d.Cache[id] = user
	}
	return user
}

func (d *CachingDecorator) GetUsersByIDs(ids []int32) []*db.User {
	// For simplicity, caching for this method is not implemented
	return d.next.GetUsersByIDs(ids)
}

func (d *CachingDecorator) SearchUsers(fname, city string, phone string, married bool) []*db.User {
	// For simplicity, caching for this method is not implemented
	return d.next.SearchUsers(fname, city, phone, married)
}
