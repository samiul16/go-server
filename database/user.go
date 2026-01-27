package database

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner string `json:"is_shop_owner"`
}

var users []User

func (usr User) Store() User {
	if usr.ID != 0 {
		return usr
	}

	users = append(users, usr)

	return usr
}

func Find(email, password string) *User {
	for _, usr := range users {
		if usr.Email == email && usr.Password == password {
			return &usr
		}
	}
	return nil
}
