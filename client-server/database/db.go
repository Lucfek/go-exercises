/*Package db is designed for working witch small databes
of users stored in a json file*/
package database

import (
	"encoding/json"
	"errors"
	mutex "go-exercises/client-server/mutex"
	"io/ioutil"
	"log"
	"os"
	"time"
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
	m     mutex.Mutex
	users map[uint64]User
	file  string
}

type saveData struct {
	LastID uint64
	Users  map[uint64]User
}

//Get gets the user with the given id
func (d Database) Get(id uint64) (User, error) {
	d.m.Lock()
	defer d.m.Unlock()
	if user, ok := d.users[id]; ok {
		return user, nil
	}
	return User{}, errors.New("There is no user with matching id")

}

//Set adds user to the database
func (d *Database) Set(user User) User {
	d.m.Lock()
	defer d.m.Unlock()
	if len(d.users) == 0 {
		user.ID = 1
	} else {
		user.ID = lastID + 1
		lastID++
	}

	d.users[user.ID] = user
	return user
}

//Delete deletes the user with the given id
func (d *Database) Delete(id uint64) (User, error) {
	d.m.Lock()
	defer d.m.Unlock()
	if user, ok := d.users[id]; ok {
		delete(d.users, id)
		return user, nil
	}

	return User{}, errors.New("There is no user with matching id")
}

//Save saves databases in file "database.json"
func (d Database) Save() error {
	d.m.Lock()
	defer d.m.Unlock()
	sd := saveData{
		LastID: lastID,
		Users:  d.users,
	}
	data, err := json.Marshal(sd)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(d.file, []byte(data), 0777)
}

//Saving saves database witch given delay, it must be run in gorutine
func (d Database) Saving(delay int) {
	for {
		err := d.Save()
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Database Saved")
		}
		time.Sleep(time.Duration(delay) * time.Minute)
	}
}

//New crates new Database
func New(file string) (*Database, error) {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			sd := saveData{Users: make(map[uint64]User)}
			data, err := json.Marshal(sd)
			if err != nil {
				return &Database{}, err
			}
			ioutil.WriteFile(file, []byte(data), 0777)
		} else {
			return &Database{}, err
		}
	}
	d := &Database{m: mutex.New(), users: map[uint64]User{}, file: file}
	return d, nil
}

//Load old Database
func (d *Database) Load() error {
	data, err := ioutil.ReadFile(d.file)
	if err != nil {
		return err
	}
	sd := saveData{Users: make(map[uint64]User)}
	err = json.Unmarshal(data, &sd)
	if err != nil {
		return err
	}
	lastID = sd.LastID
	d.users = sd.Users
	return nil
}
