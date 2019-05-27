package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var lastID uint64 = 1

//User structure of the user data
type User struct {
	ID      uint64
	Name    string
	Surname string
	Email   string
}

//Database structure of the database
type Database struct {
	m     sync.RWMutex
	users map[uint64]User
}

//Get gets the user with the given id
func (d *Database) Get(id uint64) ([]byte, error) {
	d.m.Lock()
	defer d.m.Unlock()
	if _, ok := d.users[id]; ok {
		return json.Marshal(d.users[id])
	}
	return nil, errors.New("There is no user witch matching id")

}

//Set adds user to the database
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

//Delete deletes the user with the given id
func (d *Database) Delete(id uint64) error {
	d.m.Lock()
	defer d.m.Unlock()
	if _, ok := d.users[id]; ok {
		delete(d.users, id)
		return nil
	}

	return errors.New("There is no user with matching id")
}

//Save saves databases in file "database.json"
func (d *Database) Save() error {
	d.m.Lock()
	defer d.m.Unlock()
	m := map[uint64]map[uint64]User{
		lastID: d.users,
	}
	fmt.Print()
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("database/database.json", []byte(data), 0777)
}

//New crates new Database
func New() *Database {
	return &Database{users: map[uint64]User{}}
}

//Load old Database
func Load() (*Database, error) {
	data, err := ioutil.ReadFile("database/database.json")
	if err != nil {
		return &Database{}, err
	}
	var m map[uint64]map[uint64]User
	var d *Database
	err = json.Unmarshal(data, &m)
	if err != nil {
		return &Database{}, err
	}
	for key, value := range m {
		lastID = key
		d = &Database{users: value}
	}
	return d, nil
}

//SetHandler handles POST requests
func (d *Database) SetHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	user.Name = r.FormValue("name")
	user.Surname = r.FormValue("surname")
	user.Email = r.FormValue("email")

	err := d.Set(*user)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("User successfully added"))
}

//GetHandler handles GET requests
func (d *Database) GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	user, err := d.Get(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(user)
}

//DelHandler handles DELETE requests
func (d *Database) DelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	err = d.Delete(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("User with id- " + vars["id"] + " successfully deleted"))
}
