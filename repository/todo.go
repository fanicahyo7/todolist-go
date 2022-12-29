package repository

import (
	"database/sql"
	"todolist/model"
)

type TodoListRepository interface {
	GetByUserID(userID int) ([]model.Todo, error)
	// GetByID(id int64) (*model.Todo, error)
	// Create(todoList model.Todo) (int64, error)
	// Update(todoList model.Todo) error
	// Delete(id int64) error
}

type todoListRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoListRepository {
	return &todoListRepository{db: db}
}

func (r *todoListRepository) GetByUserID(userID int) ([]model.Todo, error) {
	rows, err := r.db.Query("SELECT id, title, created_at FROM todos WHERE id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todoLists []model.Todo
	for rows.Next() {
		var todoList model.Todo
		err = rows.Scan(&todoList.ID, &todoList.Title, &todoList.Created)
		if err != nil {
			return nil, err
		}
		todoLists = append(todoLists, todoList)
	}
	return todoLists, nil
}
