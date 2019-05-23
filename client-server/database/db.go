package db

import (
	"sync"
)

//User typego
type User struct {
	ID      uint64
	Name    string
	Surname string
	Email   string
}

type database struct {
	m     sync.RWMutex
	users map[uint64]User
}

func (d *database) Get(id uint64) (User, string) {
	d.m.Lock()
	defer d.m.Unlock()
	if _, ok := d.users[id]; ok {
		return d.users[id], ""
	}
	return User{}, "There is no user witch matching id"

}

func (d *database) Set(user User) {
	d.m.Lock()
	defer d.m.Unlock()
	if len(d.users) == 0 {
		user.ID = 0
	} else {

	}

	d.users[user.ID] = user
}

func (d *database) Delete(id uint64) {
	d.m.Lock()
	defer d.m.Unlock()
	delete(d.users, id)
}

//New crates new database
func New() *database {
	mutex := new(sync.RWMutex)
	return &database{*mutex, map[uint64]User{}}
}

//Load old database
/*func Load(fileName string) database {
	//Loading from file

}*/
