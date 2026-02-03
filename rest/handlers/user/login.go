package user

import (
	"encoding/json"
	"fmt"
	"go-server/utils"
	"net/http"
)


type ReqLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login request with:", r.Body)
	var reqLogin ReqLogin
	decoded := json.NewDecoder(r.Body)
	fmt.Println("login request with:", decoded)
	err := decoded.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
	}


	usr, err := h.userRepo.Find(reqLogin.Email)
	if err != nil {
		http.Error(w, "Invalid Credential", http.StatusBadRequest)
		return
	}

	jwt, err := utils.CreateJwt("my-screate", utils.Payload{
		Email: usr.Email,
		Sub:   usr.Password,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusBadRequest)
		return
	}


	utils.SendData(w, jwt)

}