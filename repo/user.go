package repo

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

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


type UserRepo interface {
	Create(user User) (*User, error)
	Find(eamil string) (*User, error)
	// List() ([]*User, error)
	// Delete(userId int) error
	// Update(user User) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func (r *userRepo) Create(user User) (*User, error) {
	query := `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			is_shop_owner
		) VALUES (
			:first_name,
			:last_name,
			:email,
			:password,
			:is_shop_owner
		)
		RETURNING id, created_at, updated_at
	`

	var userID int
	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		fmt.Println("Error happend", err)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&userID)
	}

	user.ID= userID

	return &user, nil
}
func (r *userRepo) Find(email string) (*User, error) {
	var user User

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			password,
			is_shop_owner,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`

	err := r.db.Get(&user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}


// func (r *userRepo) List() ([]*User, error) {
// 	return r.users, nil
// }
// func (r *userRepo) Delete(productId int) error {
// 	var temList []*User

// 	for _, product := range r.users {
// 		if product.ID != productId {
// 			temList = append(temList, product)
// 		}
// 	}

// 	r.users = temList

// 	return nil
// }
// func (r *userRepo) Update(user User) (*User, error) {
// 	for idx, p := range r.users {
// 		if p.ID == user.ID {
// 			r.users[idx] = &user
// 		}
// 	}
// 	return &user, nil
// }

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
