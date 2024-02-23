package repository

import (
	"database/sql"
	"github.com/ana-flav/authentication-api.git/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository (db *sql.DB) *UserRepository{
	return &UserRepository{
		db:db,
	}
}

func (ur *UserRepository) insertUser(user models.UserModel) (models.UserModel, error) {
    var idInsertedUser models.UserModel

    err := ur.db.QueryRow("INSERT INTO user (username, password) VALUES ($1, $2) RETURNING id", user.Username, user.Password).
        Scan(&idInsertedUser.ID)
	if err != nil {
		return models.UserModel{}, err
	}
    return idInsertedUser, err
}
