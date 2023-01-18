package repository

import (
	"database/sql"
	"fmt"
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
	rows, err := r.db.Query("SELECT id, title, description, status, created_at, updated_at FROM todos WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todoLists []model.Todo
	for rows.Next() {
		var todoList model.Todo
		err = rows.Scan(&todoList.ID, &todoList.Title, &todoList.Description, &todoList.Status, &todoList.Created, &todoList.Updated)
		if err != nil {
			return todoLists, err
		}
		todoLists = append(todoLists, todoList)
	}
	fmt.Println("lanjut")
	return todoLists, nil
}
