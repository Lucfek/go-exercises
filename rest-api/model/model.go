package model

import "database/sql"

// Model struct
type Model struct {
	db *sql.DB
}

// Todo is a structure of database info
type Todo struct {
	ID          int
	Name        string
	Description string
	CratedAt    string
	UpdatedAt   string
}

// New gets address of databas as parameter  od returns new Model struct
func New(db *sql.DB) Model {
	return Model{db: db}
}

// Close ends connection with database
func (m Model) Close() {
	m.db.Close()
}

// Set inserts "todo" into database
func (m Model) Set(todo Todo) (Todo, error) {
	sqlStatement := `INSERT INTO todos (name, description) VALUES($1, $2) 
		RETURNING id, created_at, updated_at`
	err := m.db.QueryRow(sqlStatement, todo.Name, todo.Description).Scan(
		&todo.ID, &todo.CratedAt, &todo.UpdatedAt)
	return todo, err
}

// Get gets row of specified id from database
func (m Model) Get(id uint64) (Todo, error) {
	todo := Todo{}
	sqlStatement := `SELECT id, name, description, created_at, updated_at FROM todos WHERE id=$1`
	err := m.db.QueryRow(sqlStatement, id).Scan(&todo.ID, &todo.Name, &todo.Description, &todo.CratedAt, &todo.UpdatedAt)
	return todo, err
}

// GetAll gets all rows from database
func (m Model) GetAll() ([]Todo, error) {
	rows, err := m.db.Query("SELECT id, name, description, created_at, updated_at FROM todos ORDER BY created_at")
	if err != nil {
		return []Todo{}, err
	}
	var todos []Todo
	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.CratedAt, &todo.UpdatedAt)
		if err != nil {
			return []Todo{}, err
		}
		todos = append(todos, todo)
	}
	return todos, err
}

// Update updates row of specified id from database
func (m Model) Update(id uint64, todo Todo) (Todo, error) {
	sqlStatement := `UPDATE todos SET name = $1, description = $2, updated_at = now() 
		WHERE id=$3 RETURNING id, created_at, updated_at`
	err := m.db.QueryRow(sqlStatement, todo.Name, todo.Description, id).Scan(
		&todo.ID, &todo.CratedAt, &todo.UpdatedAt)
	return todo, err
}

// Delete deletes row of specified id from database
func (m Model) Delete(id uint64) (Todo, error) {
	todo := Todo{}
	sqlStatement := `DELETE FROM todos WHERE id=$1 RETURNING id, name, description, created_at, updated_at`
	err := m.db.QueryRow(sqlStatement, id).Scan(
		&todo.ID, &todo.Name, &todo.Description, &todo.CratedAt, &todo.UpdatedAt)
	return todo, err
}
