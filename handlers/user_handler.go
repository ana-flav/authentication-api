package handlers

import (
	"encoding/json"
	"go/token"
	"net/http"

	"github.com/ana-flav/authentication-api.git/models"
	"github.com/ana-flav/authentication-api.git/repository"
	"github.com/ana-flav/authentication-api.git/utils"
)


type UserHandler struct {
	userRepository repository.UserRepository 
	password utils.PasswordHash
}

func NewSongService(userRepository repository.UserRepository, password utils.PasswordHash) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
		password: password,
	}
}

func (uh *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.UserModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	if len(user.Username) < 5 && len(user.Username) > 150 {
		http.Error(w, "Username length must be between 5 and 150 characters", http.StatusBadRequest)
		return
	} else if len(user.Password) < 8 || len(user.Password) > 150{
		http.Error(w, "Password length must be between 8 and 150 characters", http.StatusBadRequest)
		return
	}

	hashedPassword, err := uh.password.HashPassword(user.Password)

	if err != nil{
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	userID, resultErr := uh.userRepository.insertUser(models.UserModel{
		Username: user.Username,
		Password: hashedPassword,
	})
	if resultErr != nil {
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return
	}

	token, err = utils.GenerateToken(userID, user.Username)

}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.UserModel
	json.NewDecoder(r.Body).Decode(&user)


}
