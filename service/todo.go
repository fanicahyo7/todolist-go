package service

import (
	"todolist/model"
	"todolist/repository"
)

type TodoListService interface {
	GetByUserID(userID int64) ([]model.Todo, error)
	// GetByID(id int64) (*model.Todo, error)
	// Create(todoList model.Todo) (int64, error)
	// Update(todoList model.Todo) error
	// Delete(id int64) error
}

type todoListService struct {
	repo repository.TodoListRepository
}

func NewTodoListService(repo repository.TodoListRepository) TodoListService {
	return &todoListService{repo: repo}
}

func (s *todoListService) GetByUserID(userID int64) ([]model.Todo, error) {
	todoLists, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return todoLists, nil
}

// func (s *todoListService) GetByID(id int64) (*model.Todo, error) {
// 	todoList, err := s.repo.GetByID(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return todoList, nil
// }

// func (s *todoListService) Create(todoList model.Todo) (int64, error) {
// 	id, err := s.repo.Create(todoList)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return id, nil
// }

// func (s *todoListService) Update(todoList model.Todo) error {
// 	err := s.repo.Update(todoList)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *todoListService) Delete(id int64) error {
// 	err := s.repo.Delete(id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
