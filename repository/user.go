package repository

import (
	"database/sql"
	"errors"
	"time"
	"todolist/model"
)

type UserRepository interface {
	GetUser(username string, password string) (model.User, error)
	RegisterUser(user model.User, passwordhHas string) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUser(usernameOrEmail, password string) (model.User, error) {
	var user model.User
	rows, err := r.db.Query("SELECT id, username FROM users WHERE (username = ? or email = ?) AND password = ?", usernameOrEmail, usernameOrEmail, password)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return user, err
		}
	} else {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (r *userRepository) RegisterUser(user model.User, passwordhHas string) (model.User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.User{}, err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO users (username, email, password) VALUES (?,?,?)", user.Username, user.Email, passwordhHas)
	if err != nil {
		return model.User{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.User{}, err
	}

	row := r.db.QueryRow("SELECT * FROM users WHERE username=? AND email=?", user.Username, user.Email)
	var createdAt []byte
	var updatedAt []byte
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &createdAt, &updatedAt)
	if err != nil {
		return model.User{}, err
	}

	user.Created, err = time.Parse("2006-01-02 15:04:05.999999", string(createdAt))
	if err != nil {
		return model.User{}, err
	}
	user.Updated, err = time.Parse("2006-01-02 15:04:05.999999", string(updatedAt))
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// func (r *userRepository) RegisterUsers(users []model.User) error {
// 	tx, err := r.db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()

// 	for _, user := range users {
// 		// hash the password using bcrypt
// 		hashedPassword, err := hashPassword(user.Password)
// 		if err != nil {
// 			return err
// 		}

// 		// insert the user into the database
// 		_, err = tx.Exec("INSERT INTO users (username, email, password) VALUES (?,?,?)", user.Username, user.Email, hashedPassword)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	if err := tx.Commit(); err != nil {
// 		return err
// 	}
// 	return nil
// }
