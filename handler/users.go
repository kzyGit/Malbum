package handler

import (
	"encoding/json"
	"net/http"
)

type Users struct {
    ID int `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}

var users []Users

type usersHandler struct{}

func (handler *usersHandler) userSignup(w http.ResponseWriter, req *http.Request) {
	var user Users
	_ = json.NewDecoder(req.Body).Decode(&user)
	user.ID = id + 1
	id++
	users = append(users, user)
	json.NewEncoder(w).Encode(user)

}
