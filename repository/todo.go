package repository

import (
	"database/sql"
	"time"
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
	rows, err := r.db.Query("SELECT id, user_id, title, description, status, created_at FROM todos WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todoLists []model.Todo
	for rows.Next() {
		var todoList model.Todo
		var createdAt []byte
		err = rows.Scan(&todoList.ID, &todoList.UserID, &todoList.Title, &todoList.Description, &todoList.Status, &createdAt)
		if err != nil {
			return todoLists, err
		}
		todoList.Created, err = time.Parse("2006-01-02 15:04:05.999999", string(createdAt))
		if err != nil {
			return todoLists, err
		}

		todoLists = append(todoLists, todoList)
	}
	return todoLists, nil
}
