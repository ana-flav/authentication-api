package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/ana-flav/authentication-api.git/models"
	"github.com/ana-flav/authentication-api.git/repository"
)


type UserHandler struct {
	userRepository repository.UserRepository 
}

func NewSongService(userRepository repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
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


}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.UserModel
	json.NewDecoder(r.Body).Decode(&user)


}
