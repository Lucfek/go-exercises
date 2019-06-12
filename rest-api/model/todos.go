package model

import "database/sql"

// Todos struct
type Todos struct {
	DB *sql.DB
}

// Todo is a structure of database info
type Todo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CratedAt    string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// Set inserts "todo" into database
func (m Todos) Set(todo Todo) (Todo, error) {
	sqlStatement := `INSERT INTO todos (name, description) VALUES($1, $2) 
		RETURNING id, created_at, updated_at`
	err := m.DB.QueryRow(sqlStatement, todo.Name, todo.Description).Scan(
		&todo.ID, &todo.CratedAt, &todo.UpdatedAt)
	return todo, err
}

// Get gets row of specified id from database
func (m Todos) Get(id uint64) (Todo, error) {
	todo := Todo{}
	sqlStatement := `SELECT id, name, description, created_at, updated_at FROM todos WHERE id=$1`
	err := m.DB.QueryRow(sqlStatement, id).Scan(&todo.ID, &todo.Name, &todo.Description, &todo.CratedAt, &todo.UpdatedAt)
	return todo, err
}

// GetAll gets all rows from database
func (m Todos) GetAll() ([]Todo, error) {
	rows, err := m.DB.Query("SELECT id, name, description, created_at, updated_at FROM todos ORDER BY created_at")
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
	return todos, rows.Err()
}

// Update updates row of specified id from database
func (m Todos) Update(id uint64, todo Todo) (Todo, error) {
	sqlStatement := `UPDATE todos SET name = $1, description = $2, updated_at = now() 
		WHERE id=$3 RETURNING id, created_at, updated_at`
	err := m.DB.QueryRow(sqlStatement, todo.Name, todo.Description, id).Scan(
		&todo.ID, &todo.CratedAt, &todo.UpdatedAt)
	return todo, err
}

// Delete deletes row of specified id from database
func (m Todos) Delete(id uint64) (Todo, error) {
	todo := Todo{}
	sqlStatement := `DELETE FROM todos WHERE id=$1 RETURNING id, name, description, created_at, updated_at`
	err := m.DB.QueryRow(sqlStatement, id).Scan(
		&todo.ID, &todo.Name, &todo.Description, &todo.CratedAt, &todo.UpdatedAt)
	return todo, err
}
