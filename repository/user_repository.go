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

func (ur *UserRepository) insertUser(user models.UserModel) (error) {
	_,err := ur.db.Exec("INSERT INTO user (user, senha) values ($1, $2)", user.Username, user.Password)
	return  err
}