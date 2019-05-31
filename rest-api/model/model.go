package model

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Model struct {
	db *sql.DB
}

type Todo struct {
	Id          int    `json:"Id,sting"`
	Title       string `json:"Title"`
	Description string `json:"Desctription"`
	Crated_at   string `"json:Created_at"`
	Updated_at  string `"json:Updated_at"`
}

func New(dbAddr string) (*Model, error) {
	db, err := sql.Open("postgres", dbAddr)
	if err != nil {
		return &Model{}, err
	}
	return &Model{db: db}, nil
}
func (m Model) Close() {
	m.db.Close()
}

func (m Model) Set(todo Todo) (Todo, error) {
	return Todo{Id: 1, Title: "test", Description: "test", Crated_at: "test", Updated_at: "test"}, nil
}
func (m Model) GetAll() (Todo, error) {
	return Todo{Id: 1, Title: "test", Description: "test", Crated_at: "test", Updated_at: "test"}, nil
}
func (m Model) Get(id string) (Todo, error) {
	return Todo{Id: 1, Title: id, Description: "test", Crated_at: "test", Updated_at: "test"}, nil
}
func (m Model) Update(id string) (Todo, error) {
	return Todo{Id: 1, Title: id, Description: "test", Crated_at: "test", Updated_at: "test"}, nil
}
func (m Model) Delete(id string) (Todo, error) {
	return Todo{Id: 1, Title: id, Description: "test", Crated_at: "test", Updated_at: "test"}, nil
}
