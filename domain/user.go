package domain

import "time"

type User struct {
	ID          int       `db:"id" json:"id"`
	FirstName   string    `db:"first_name" json:"first_name"`
	LastName    string    `db:"last_name" json:"last_name"`
	Email       string    `db:"email" json:"email"`
	Password    string    `db:"password" json:"password"`
	IsShopOwner bool      `db:"is_shop_owner" json:"is_shop_owner"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}