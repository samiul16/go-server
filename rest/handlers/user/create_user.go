package user

import (
	"encoding/json"
	"fmt"
	"go-server/repo"
	"go-server/utils"
	"net/http"
)

type ReqUser struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request)  {
	var newUser ReqUser
	decoded := json.NewDecoder(r.Body)
	err := decoded.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
	}

	fmt.Println("New User",newUser)

	createdUser, err := h.userRepo.Create(repo.User{
		FirstName: newUser.FirstName,
		LastName: newUser.LastName,
		Email: newUser.Email,
		Password: newUser.Password,
		IsShopOwner: newUser.IsShopOwner,
	})

	utils.SendData(w, createdUser)

}

