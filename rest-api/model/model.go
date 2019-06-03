package model

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq" //Database driver
)

type Model struct {
	db *sql.DB
}

type Todo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CratedAt    string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
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
	sqlStatement := `INSERT INTO todos (name, description) VALUES($1, $2) RETURNING id, created_at, updated_at;`
	err := m.db.QueryRow(sqlStatement, todo.Name, todo.Description).Scan(
		&todo.Id, &todo.CratedAt, &todo.UpdatedAt)
	return todo, err
}
func (m Model) Get(id uint64) (Todo, error) {
	todo := Todo{}
	sqlStatement := `SELECT id, name, description, created_at, updated_at FROM todos WHERE id=$1;`
	row := m.db.QueryRow(sqlStatement, id)
	err := row.Scan(&todo.Id, &todo.Name, &todo.Description, &todo.CratedAt, &todo.UpdatedAt)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return Todo{}, errors.New("Zero rows found")
		default:
			return Todo{}, err
		}
	}
	return todo, nil
}
func (m Model) GetAll() ([]Todo, error) {
	rows, err := m.db.Query("SELECT id, name, description, created_at, updated_at FROM todos;")
	if err != nil {
		return []Todo{}, err
	}
	var todos []Todo
	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.Id, &todo.Name, &todo.Description, &todo.CratedAt, &todo.UpdatedAt)
		if err != nil {
			return []Todo{}, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
func (m Model) Update(id uint64, todo Todo) (Todo, error) {
	sqlStatement := `UPDATE todos SET name = $1, description = $2, updated_at = now() WHERE id=$3 RETURNING id, created_at, updated_at;`
	err := m.db.QueryRow(sqlStatement, todo.Name, todo.Description, id).Scan(
		&todo.Id, &todo.CratedAt, &todo.UpdatedAt)
	return todo, err
}
func (m Model) Delete(id uint64) (Todo, error) {
	return Todo{}, nil
}
