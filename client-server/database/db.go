package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"
)

var lastID uint64 = 1

//User typego
type User struct {
	ID      uint64
	Name    string
	Surname string
	Email   string
}

type Database struct {
	m     sync.RWMutex
	users map[uint64]User
}

func (d *Database) Get(id uint64) (User, error) {
	d.m.Lock()
	defer d.m.Unlock()
	if _, ok := d.users[id]; ok {
		return d.users[id], nil
	}
	return User{}, errors.New("There is no user witch matching id")

}

func (d *Database) Set(user User) error {
	d.m.Lock()
	defer d.m.Unlock()
	if len(d.users) == 0 {
		user.ID = 1
	} else {
		user.ID = lastID + 1
		lastID++
	}

	d.users[user.ID] = user
	if _, ok := d.users[lastID]; ok {
		return nil
	}
	return errors.New("Adding user faild")
}

func (d *Database) Delete(id uint64) error {
	d.m.Lock()
	defer d.m.Unlock()
	if _, ok := d.users[id]; ok {
		delete(d.users, id)
		return nil
	}

	return errors.New("There is no user witch matching id")
}

func (d *Database) Save() {
	d.m.Lock()
	defer d.m.Unlock()
	m := map[uint64]map[uint64]User{
		lastID: d.users,
	}
	fmt.Print()
	data, _ := json.Marshal(m)
	ioutil.WriteFile("database/database.json", []byte(data), 0777)
}

//New crates new Database
func New() *Database {
	return &Database{users: map[uint64]User{}}
}

//Load old Database
func Load() *Database {
	data, _ := ioutil.ReadFile("database/database.json")
	var m map[uint64]map[uint64]User
	var d *Database
	json.Unmarshal(data, &m)
	for key, value := range m {
		lastID = key
		d = &Database{users: value}
	}
	return d
}
