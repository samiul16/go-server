package handlers

import (
	"encoding/json"
	"fmt"
	"go-server/database"
	"go-server/utils"
	"net/http"
)


type ReqLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	decoded := json.NewDecoder(r.Body)
	err := decoded.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
	}


	usr := database.Find(reqLogin.Email, reqLogin.Password)

	if usr == nil {
		http.Error(w, "Invalid Credential", http.StatusBadRequest)
	}


	utils.SendData(w, usr)

}