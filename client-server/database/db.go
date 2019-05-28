package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

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
type Mutex chan struct{}

func (m Mutex) Lock() {
	<-m
}

func (m Mutex) Unlock() {
	m <- struct{}{}
}

//Database structure of the database
type Database struct {
	m     Mutex
	users map[uint64]User
}

//Get gets the user with the given id
func (d *Database) Get(id uint64) ([]byte, error) {
	d.m.Lock()
	defer d.m.Unlock()
	if _, ok := d.users[id]; ok {
		return json.Marshal(d.users[id])
	}
	return nil, errors.New("There is no user with matching id")

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
	d := &Database{m: make(Mutex, 1), users: map[uint64]User{}}
	d.m <- struct{}{}
	return d
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
		d = &Database{m: make(Mutex, 1), users: value}
	}
	d.m <- struct{}{}
	return d, nil
}

//SetHandler handles POST requests
func (d *Database) SetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	user := new(User)
	user.Name = r.FormValue("name")
	user.Surname = r.FormValue("surname")
	user.Email = r.FormValue("email")

	err := d.Set(*user)
	if err != nil {
		io.WriteString(w, `{"Error":"`+err.Error()+`"}`)
		return
	}
	io.WriteString(w, `{"Status":"Added Succesfully", "UserId":`+strconv.FormatUint(lastID, 10)+`}`)
}

//GetHandler handles GET requests
func (d *Database) GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		io.WriteString(w, `{"Error":"`+err.Error()+`"}`)
		return
	}
	user, err := d.Get(id)
	if err != nil {
		io.WriteString(w, `{"Error":"`+err.Error()+`"}`)
		return
	}
	io.WriteString(w, string(user))
}

//DelHandler handles DELETE requests
func (d *Database) DelHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		io.WriteString(w, `{"Error":"`+err.Error()+`"}`)
		return
	}
	err = d.Delete(id)
	if err != nil {
		io.WriteString(w, `{"Error":"`+err.Error()+`"}`)
		return
	}
	io.WriteString(w, `{"Status":"Deleted User[`+vars["id"]+`] Succesfully"}`)
}
