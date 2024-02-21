package models

import (
	"github.com/google/uuid"
)

type UserModel struct {
	ID uuid.UUID `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}