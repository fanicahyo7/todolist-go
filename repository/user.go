package repository

import (
	"database/sql"
	"errors"
	"time"
	"todolist/model"
)

type UserRepository interface {
	GetByUserID(userID int) (model.User, error)
	GetUserWithPassword(usernameOrEmail, password string) (model.User, error)
	GetUser(usernameOrEmail string) (model.User, error)
	SaveUser(user model.User, passwordhHas string) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByUserID(userID int) (model.User, error) {
	var user model.User
	var createdAt []byte
	var updatedAt []byte
	rows, err := r.db.Query("SELECT id, username, password, email, created_at, updated_at FROM users WHERE id = ?", userID)
	if err != nil {
		return model.User{}, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &createdAt, &updatedAt); err != nil {
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
	} else {
		return model.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r *userRepository) GetUserWithPassword(usernameOrEmail, password string) (model.User, error) {
	var user model.User
	rows, err := r.db.Query("SELECT id, username FROM users WHERE (username = ? or email = ?) AND password = ?", usernameOrEmail, usernameOrEmail, password)
	if err != nil {
		return model.User{}, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return model.User{}, err
		}

	} else {
		return model.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r *userRepository) GetUser(usernameOrEmail string) (model.User, error) {
	var user model.User
	rows, err := r.db.Query("SELECT id, username, password FROM users WHERE username = ? or email = ? ", usernameOrEmail, usernameOrEmail)
	if err != nil {
		return model.User{}, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return model.User{}, err
		}

	} else {
		return model.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r *userRepository) SaveUser(user model.User, passwordhHas string) (model.User, error) {
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

// func (r *userRepository) LoginUser(user model.User, passwordhHas string) (model.User, error) {

// }

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
