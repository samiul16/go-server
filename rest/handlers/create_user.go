package handlers

import (
	"encoding/json"
	"fmt"
	"go-server/database"
	"go-server/utils"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request)  {
	var newUser database.User
	decoded := json.NewDecoder(r.Body)
	err := decoded.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
	}

	createdUser := newUser.Store()

	utils.SendData(w, createdUser)

}

